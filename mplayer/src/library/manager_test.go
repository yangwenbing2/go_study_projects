package library

import "testing"

func TestManager(t *testing.T) {
	// 测试创建
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	// 测试长度
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	// 测试增加歌曲
	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celion Dion",
		"http://qbox.me/24501234", "MP3", "music"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	// 测试查找
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}

	//
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	//测试删除
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}
}
