package main

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

type Tree struct {
	leftVal, rightVal *Tree
	curVal            int
}

func (t *Tree) insert(v int) (err error) {

	if t == nil {
		return errors.New("Cannot insert in empty tree")
	}

	if v < t.curVal {

		if t.leftVal == nil {
			log.Info("Current value less than <Tree.curVal>. Insert new child-left-tree")
			t.leftVal = &Tree{
				curVal: v,
			}
			return
		}

		log.Info("child-left-tree not empty, insert current value to left tree")
		return t.leftVal.insert(v)
	}

	if v > t.curVal {
		log.Info("Current value greater than <Tree.curVal>. Insert new child-right-tree")
		if t.rightVal == nil {
			t.rightVal = &Tree{
				curVal: v,
			}
			return
		}

		log.Info("child-right-tree not empty, insert current value to right tree")
		return t.rightVal.insert(v)
	}

	log.Info("This value equal <Tree.curVal>. Do nothing")
	return
}

func (t *Tree) search(v int) (res bool, err error) {
	if t == nil {
		log.Info("Tree is empty")
		return false, errors.New("Cannot search in empty tree")
	}

	if v < t.curVal {
		log.Info("Current value less than <Tree.curVal>. Search value in child-left-tree")
		return t.leftVal.search(v)
	}

	if v > t.curVal {
		log.Info("Current value greater than <Tree.curVal>. Search value in child-right-tree")
		return t.rightVal.search(v)
	}

	if v == t.curVal {
		log.Info("Current value equal <Tree.curVal>. Search is end, return <true>")
		return true, nil
	}

	log.Info("Unknown error")
	return false, errors.New("Unknown error")
}

func (t *Tree) delete(v int, parent *Tree) (res bool, err error) {
	if t == nil {
		log.Info("Tree is empty")
		return false, errors.New("Cannot delete value from empty tree")
	}

	if v < t.curVal {
		log.Info("Current value less than <Tree.curVal>. Delete value in child-left-tree")
		return t.leftVal.delete(v, t)
	}

	if v > t.curVal {
		log.Info("Current value greater than <Tree.curVal>. Delete value in child-right-tree")
		return t.rightVal.delete(v, t)
	}

	log.Info("Current value equal <Tree.curVal>. Delete value")
	if t.leftVal == nil && t.rightVal == nil {
		log.Info("Tree Left and Right is empty - delete current value")

		return parent.replaceBranch(t, nil)
	}

	if t.leftVal == nil {
		log.Info("Tree Left is empty - replace current value on Right Branch")
		// t = t.rightVal
		return parent.replaceBranch(t, t.rightVal)

	}

	if t.rightVal == nil {
		log.Info("Tree Right is empty - replace current value on Left Branch")
		// t = t.leftVal
		return parent.replaceBranch(t, t.leftVal)
	}

	log.Info("Tree Right and Left is not empty - find minimal value in Right Branch, then replace Curval and delete min")
	minLeaf := t.rightVal.popMinLeaf(t)

	t.curVal = minLeaf

	return true, nil
}

func (t *Tree) replaceBranch(child *Tree, newTree *Tree) (res bool, err error) {
	if t == nil && newTree == nil {
		return false, errors.New("Cannot Delete Root Tree")
	}

	if t == nil {
		log.Info("Parent is empty. This is a root tree")
		child.curVal = newTree.curVal
		return child.replaceBranch(newTree, nil)
	}

	if t.leftVal == child {
		log.Info("Replace Left link Tree")
		t.leftVal = newTree
		return true, nil
	}

	if t.rightVal == child {
		log.Info("Replace Right link Tree")
		t.rightVal = newTree
		return true, nil
	}

	return false, errors.New("Cannot Replace child Branch")
}

func (t *Tree) popMinLeaf(parent *Tree) (min int) {
	log.Info("Get minimal leaf value and delete link on this tree")

	min = t.curVal
	log.Info(min)
	if t != nil && t.leftVal != nil {
		return t.leftVal.popMinLeaf(t)
	}

	parent.replaceBranch(t, t.rightVal)
	return min
}

func initTree(dataPath string) (dataTree *Tree, err error) {
	var data [30]int

	log.Info("Open json init file")
	file, err := os.Open(dataPath)
	if err != nil {
		return
	}

	log.Info("Decode json init file")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return
	}

	log.Info("Init root Tree with first element")
	dataTree = &Tree{
		curVal: data[0],
	}

	log.Info("Init root Tree")
	for i := 1; i < len(data); i++ {

		err = dataTree.insert(data[i])
		if err != nil {
			return
		}
	}

	return
}
