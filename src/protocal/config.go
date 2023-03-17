package protocal

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type Config struct {
	Id         int
	Key        string
	Value      string
	ClusterId  int
	Status     int8
	Version    string
	CreateTime string
	UpdateTIme string
}

func NewConfig(id int, key string, value string, clusterId int, status int8, version string) *Config {
	return &Config{
		Id:         id,
		Key:        key,
		Value:      value,
		ClusterId:  clusterId,
		Status:     status,
		Version:    version,
		CreateTime: time.Now().Format(timeFormat),
		UpdateTIme: time.Now().Format(timeFormat),
	}
}

type ConfigNode struct {
	Config Config
	Prev   *ConfigNode
	Next   *ConfigNode
}

type VersionList struct {
	Head *ConfigNode
	Tail *ConfigNode
	len  int
}

func NewVersionList() *VersionList {
	head := &ConfigNode{Config{}, nil, nil}
	tail := &ConfigNode{Config{}, nil, nil}
	head.Next = tail
	tail.Prev = head
	return &VersionList{
		Head: head,
		Tail: tail,
		len:  0,
	}
}

func (list *VersionList) Length() int {
	return list.len
}

func (list *VersionList) AddFirst(config Config) {
	configNode := &ConfigNode{config, list.Head, list.Head.Next}
	list.Head.Next.Prev = configNode
	list.Head.Next = configNode
	list.len++
}

func (list *VersionList) AddLast(config *Config) {
	configNode := &ConfigNode{*config, list.Tail.Prev, list.Tail}
	list.Tail.Prev.Next = configNode
	list.Tail.Prev = configNode
	list.len++
}

func (list *VersionList) Remove(configNode *ConfigNode) {
	configNode.Prev.Next = configNode.Next
	configNode.Next.Prev = configNode.Prev
	configNode.Next = nil
	configNode.Prev = nil
	list.len--
}

func (list *VersionList) Traverse() {
	p := list.Head.Next
	for p != list.Tail {
		fmt.Println(p.Config)
		p = p.Next
	}
}
