package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
)

type ImageContainer struct {
	repo   string
	branch string
	images map[string][]string
}

func NewImageContainer(repo, branch string) *ImageContainer {
	return &ImageContainer{
		repo:   repo,
		branch: branch,
	}
}

func (s *ImageContainer) LastHash() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/repos/%s/branches/%s", GITHUB_API, s.repo, s.branch))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Commit struct {
			Commit struct {
				Tree struct {
					Sha string `json:"sha"`
				} `json:"tree"`
			} `json:"commit"`
		} `json:"commit"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Commit.Commit.Tree.Sha, nil
}

func isImageFile(path string) bool {
	return slices.Contains([]string{".png", ".jpg", ".jpeg", ".webp"}, strings.ToLower(filepath.Ext(path)))
}

func firstDir(path string) string {
	if i := strings.Index(path, "/"); i >= 0 {
		return path[:i]
	}
	return path
}

func (s *ImageContainer) QueryImages(category, hash string) ([]string, error) {
	s.images = make(map[string][]string) // Clear previous images
	resp, err := http.Get(fmt.Sprintf("%s/repos/%s/git/trees/%s?recursive=1", GITHUB_API, s.repo, hash))
	if err != nil {
		fmt.Println("Error fetching images:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Sha  string `json:"sha"`
		Tree []struct {
			Path string `json:"path"`
		} `json:"tree"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, file := range data.Tree {
		fpath := file.Path
		subdir := firstDir(fpath)
		if isImageFile(fpath) {
			url := fmt.Sprintf("%s/gh/%s@latest/%s", JSDELIVR_CDN, s.repo, fpath)
			s.images[subdir] = append(s.images[subdir], url)
		}
	}
	return s.images[category], nil
}

func (s *ImageContainer) RandomImage(category string) (string, error) {
	lastHash, err := s.LastHash()
	if err != nil {
		return "", err
	}

	images, err := s.QueryImages(category, lastHash)
	if err != nil {
		return "", err
	}

	if len(images) == 0 {
		return "", errors.New("No image found")
	}
	return images[rand.Intn(len(images))], nil
}
