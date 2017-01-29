package rbtree

import(
  "errors"
  "fmt"
)

const RED   bool = true
const BLACK bool = false

type Rbtree struct{
  root  *node
  size  int
  cmp   func(i,j interface{})bool
}

type node struct{
  key     interface{}
  data    interface{}
  color   bool
  parent  *node
  left    *node
  right   *node
}

//New() will return a pointer of Rbtree with compare function you write.
func New(cmp func(i,j interface{})bool)*Rbtree{
  t:=new(Rbtree)
  t.cmp=cmp
  t.root=nil
  t.size=0
  return t
}

//Insert a node to the Rbtree.
func (t *Rbtree) Insert(key, data interface{}){
  t.size++
  //make a new node
  n:=new(node)
  n.key=key
  n.data=data
  n.color=RED
  n.parent=nil
  n.left=nil
  n.right=nil

  //find position and set
  if t.root==nil {
    t.root=n
  }else{
    tmpn:=t.root
    for true {
      if t.cmp(n.key,tmpn.key) {
        //left subtree
        if tmpn.left==nil {
          tmpn.left=n
          n.parent=tmpn
          break
        }else{
          tmpn=tmpn.left
        }
      }else{
        //right subtree
        if tmpn.right==nil {
          tmpn.right=n
          n.parent=tmpn
          break
        }else{
          tmpn=tmpn.right
        }
      }
    }
  }
  //rebalance
  t.insertFix(n)
  t.Inorder(t.root)
  fmt.Println("+",t.root.data)
}

func (t *Rbtree) insertFix(n *node){
  if n==t.root {
    //n is root
    n.color=BLACK
  }else if n.parent.color==RED {
    //parent is red
    if n.parent.parent.left==n.parent {
      //n's parent is left subtree of n's grandparent
      if n.parent.parent.right==nil {
        //case2: uncle is nil (=BLACK)
        if n.parent.left==n {
          //case2-1: n is left of n's parent
          n=n.parent.parent
          t.rightRotate(n)
          n.color=RED
          n.parent.color=BLACK
        }else{
          //case2-2: n is right of n's parent
          n.parent.color=BLACK
          n.parent.parent.color=RED
          t.leftRotate(n.parent.parent)
          t.insertFix(n.right)
        }
      }else if n.parent.parent.right.color==RED {
        //case1: uncle is red
        n.parent.color=BLACK
        n.parent.parent.right.color=BLACK
        n.parent.parent.color=RED
        n=n.parent.parent
        t.insertFix(n)
      }else{
        //case3: uncle is BLACK (=nil)
        if n.parent.left==n {
          //case3-1: n is left of n's parent
          n=n.parent.parent
          t.rightRotate(n)
          n.color=RED
          n.parent.color=BLACK
        }else{
          //case3-2: n is right of n's parent
          n.parent.color=BLACK
          n.parent.parent.color=RED
          t.leftRotate(n.parent.parent)
          t.insertFix(n.right)
        }
      }
    }else{
      //n's parent is right subtree of n's grandparent
      if n.parent.parent.left==nil {
        //case2: uncle is nil (=BLACK)
        if n.parent.right==n {
          //case2-1: n is right of n's parent
          n=n.parent.parent
          t.leftRotate(n)
          n.color=RED
          n.parent.color=BLACK
        }else{
          //case2-2: n is left of n's parent
          n.parent.color=BLACK
          n.parent.parent.color=RED
          t.rightRotate(n.parent.parent)
          t.insertFix(n.left)
        }
      }else if n.parent.parent.left.color==RED {
        //case1: uncle is red
        n.parent.color=BLACK
        n.parent.parent.left.color=BLACK
        n.parent.parent.color=RED
        n=n.parent.parent
        t.insertFix(n)
      }else{
        //case3: uncle is BLACK (=nil)
        if n.parent.right==n {
          //case3-1: n is right of n's parent
          n=n.parent.parent
          t.leftRotate(n)
          n.color=RED
          n.parent.color=BLACK
        }else{
          //case3-2: n is left of n's parent
          n.parent.color=BLACK
          n.parent.parent.color=RED
          t.rightRotate(n.parent.parent)
          t.insertFix(n.left)
        }
      }
    }
  }

}

func (t *Rbtree) leftRotate(n *node){
  if t.root==n {
    t.root=n.right
  }else{
    if n.parent.left==n {
      n.parent.left=n.right
    }else{
      n.parent.right=n.right
    }
  }
  n.right.parent=n.parent
  n.parent=n.right
  n.right=n.parent.left
  n.parent.left=n
}

func (t *Rbtree) rightRotate(n *node){
  if t.root==n {
    t.root=n.left
  }else{
    if n.parent.left==n {
      n.parent.left=n.left
    }else{
      n.parent.right=n.left
    }
  }
  n.left.parent=n.parent
  n.parent=n.left
  n.left=n.parent.right
  n.parent.right=n
}

//Inorder traversal
func (t *Rbtree) Inorder(n *node){
  if t.root==nil{
    return
  }
  if n==nil {
    return
  }
  t.Inorder(n.left)
  fmt.Println(n.data)
  t.Inorder(n.right)
}

//Delete a node from the Rbtree.
func (t *Rbtree) Delete(key int)error{
  return nil
}

//find data with specific key
func (t *Rbtree) Find(key interface{})(interface{},error){
  tmpn:=t.root
  for tmpn!=nil {
    if t.cmp(key,tmpn.key){
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

//clean the tree
func (t *Rbtree) Clear(){
  t.size=0
  t.root=nil
}
