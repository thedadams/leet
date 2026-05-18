// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn is_unival_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(r) = root {
            let b = r.borrow();
        return Self::is_unival(&b.left, b.val) && Self::is_unival(&b.right, b.val);
        }
        true
    }

    fn is_unival(root: &Option<Rc<RefCell<TreeNode>>>, val: i32) -> bool {
        if let Some(r) = root {
            let b = r.borrow();
            if b.val != val {
                return false;
            }

            if !Self::is_unival(&b.right, val) || !Self::is_unival(&b.left, val) {
                return false;
            }
        }

        true
    }
}
