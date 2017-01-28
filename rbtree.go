package rbtree

import(
  "errors"
)

const RED   bool = true
const BLACK bool = false
const LEFT  bool = true
const RIGHT bool = false

type Rbtree struct{
  root  *Node
  cmp   func(i,j interface{})bool
}

type node struct{
  key     interface{}
  data    interface{}
  color   bool
  subtree bool
  parent  *Node
  left    *Node
  right   *Node
}

//New() will return a pointer of Rbtree with compare function you write,
func New(cmp func(i,j interface{})bool)*Rbtree{
  t:=new(Rbtree)
  t.root=nil
  return t
}

//Insert a node to the Rbtree.
func (t *Rbtree) Insert(key, data interface{})error{
  //make a new node
  n:=new(node)
  n.key=key
  n.data=data
  n.color=RED
  n.left=nil
  n.right=nil

  //find position and set
  if t.root==nil {
    t.root=n
  }else{
    tmpn:=t.root
    for true {
      if cmp(n.key,tmpn.key) {
        //left subtree
        if tmpn.left==nil {
          tmpn.left=n
          n.parent=tmpn
          n.subtree=LEFT
          break
        }else{
          tmpn=tmpn.left
        }
      }else{
        //right subtree
        if tmpn.right==nil {
          tmpn.right=n
          n.parent=tmpn
          n.subtree=RIGHT
          break
        }else{
          tmpn=tmpn.right
        }
      }
    }
  }

  //rebalance
  
}

//Delete a node from the Rbtree.
func (t *Rbtree) Delete(key int)error{
}

//find data with specific key
func (t *Rbtree) Find(key interface{})interface{},error{
  tmpn:=t.root
  for tmpn!=nil {
    if cmp(key,tmpn.key){
      //left subtree
      tmpn=tmpn.left
    }else if key==tmpn.key {
      //equal
      return tmpn.data, nil
    }else{
      //right subtree
      tmpn=tmpn.right
    }
  }

  var out interface{}
  err:=errors.New("No such node.")
  return out,err
}
