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
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        let (px, dx, py, dy) = Self::find_depth(None, root, &x, &y, 0);
        px != py && dx == dy
    }

    fn find_depth(parent: Option<Rc<RefCell<TreeNode>>>, root: Option<Rc<RefCell<TreeNode>>>, x: &i32, y: &i32, depth: i32) -> (i32, i32, i32, i32) {
        if let Some(r) = root {
            let n = r.borrow();
            if n.val == *x {
                if let Some(p) = parent {
                    return (p.borrow().val, depth, 0, 0);
                }
                return (0, depth, 0, 0);
            }
            if n.val == *y {
                if let Some(p) = parent {
                    return (0, 0, p.borrow().val, depth);
                }
                return (0, 0, 0, depth);
            }

            let (mut plx, mut dlx, mut ply, mut dly) = Self::find_depth(Some(r.clone()), n.left.clone(), x, y, depth + 1);
            let (prx, drx, pry, dry) = Self::find_depth(Some(r.clone()), n.right.clone(), x, y, depth + 1);

            if plx == 0 {
                plx = prx;
            }
            if dlx == 0 {
                dlx = drx;
            }
            if ply == 0 {
                ply = pry;
            }
            if dly == 0 {
                dly = dry;
            }

            return (plx, dlx, ply, dly);
        }

        (0, 0, 0, 0)
    }
}
