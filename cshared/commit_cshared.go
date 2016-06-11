// +build ignore
package main

import (
	"C"
	"gopkg.in/src-d/go-git.v3/core"
	. "github.com/src-d/go-git"
)

//export c_Commit
func c_Commit() uint64 {
	commit := Commit{}
	handle := RegisterObject(commit)
	return uint64(handle)
}

//export c_Commit_get_Hash
func c_Commit_get_Hash(c uint64) []byte {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return nil
	}
	commit := obj.(Commit)
	return commit.Hash[:]
}

//export c_Commit_get_Author
func c_Commit_get_Author(c uint64) uint64 {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return IH
	}
	commit := obj.(Commit)
	author := commit.Author
	author_handle := RegisterObject(author)
	return uint64(author_handle)
}

//export c_Commit_get_Committer
func c_Commit_get_Committer(c uint64) uint64 {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return IH
	}
	commit := obj.(Commit)
	committer := commit.Committer
	committer_handle := RegisterObject(committer)
	return uint64(committer_handle)
}

//export c_Commit_get_Message
func c_Commit_get_Message(c uint64) string {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return ""
	}
	commit := obj.(Commit)
	return commit.Message
}

//export c_Commit_Tree
func c_Commit_Tree(c uint64) uint64 {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return IH
	}
	commit := obj.(Commit)
	tree := commit.Tree()
	tree_handle := RegisterObject(tree)
	return uint64(tree_handle)
}

//export c_Commit_Parents
func c_Commit_Parents(c uint64) uint64 {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return IH
	}
	commit := obj.(Commit)
	parents := commit.Parents()
	parents_handle := RegisterObject(parents)
	return uint64(parents_handle)
}

//export c_Commit_NumParents
func c_Commit_NumParents(c uint64) int {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return -1
	}
	commit := obj.(Commit)
	return commit.NumParents()
}

//export c_Commit_File
func c_Commit_File(c uint64, path string) (uint64, int, string) {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return IH, ErrorCodeNotFound, MessageNotFound
	}
	commit := obj.(Commit)
	file, err := commit.File(CopyString(path))
	if err != nil {
		return IH, ErrorCodeInternal, err.Error()
	}
	file_handle := RegisterObject(file)
	return uint64(file_handle), ErrorCodeSuccess, ""
}

//export c_Commit_ID
func c_Commit_ID(c uint64) []byte {
	return c_Commit_get_Hash(c)
}

//export c_Commit_Type
func c_Commit_Type(c uint64) int8 {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return -1
	}
	commit := obj.(Commit)
	return int8(commit.Type())
}

//export c_Commit_Decode
func c_Commit_Decode(c uint64, o uint64) (int, string) {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return ErrorCodeNotFound, MessageNotFound
	}
	commit := obj.(Commit)
	obj, ok = GetObject(Handle(o))
	if !ok {
		return ErrorCodeNotFound, MessageNotFound
	}
	cobj := obj.(core.Object)
	err := commit.Decode(cobj)
	if err == nil {
		return ErrorCodeSuccess, ""
	}
	return ErrorCodeInternal, err.Error()
}

//export c_Commit_String
func c_Commit_String(c uint64) string {
	obj, ok := GetObject(Handle(c))
	if !ok {
		return ""
	}
	commit := obj.(Commit)
	return commit.String()
}

//export c_NewCommitIter
func c_NewCommitIter(r uint64, iter uint64) uint64 {
	obj, ok := GetObject(Handle(r))
	if !ok {
		return IH
	}
	repo := obj.(Repository)
	obj, ok = GetObject(Handle(iter))
	if !ok {
		return IH
	}
	obj_iter := obj.(core.ObjectIter)
	commit_iter := NewCommitIter(&repo, obj_iter)
	handle := RegisterObject(commit_iter)
	return uint64(handle)
}

//export c_CommitIter_Next
func c_CommitIter_Next(iter uint64) (uint64, int, string) {
	obj, ok := GetObject(Handle(iter))
	if !ok {
		return IH, ErrorCodeNotFound, MessageNotFound
	}
	commitIter := obj.(CommitIter)
	commit, err := commitIter.Next()
	if err != nil {
		return IH, ErrorCodeInternal, err.Error()
	}
	handle := RegisterObject(commit)
	return uint64(handle), ErrorCodeSuccess, ""
}
