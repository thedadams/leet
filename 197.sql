# Write your MySQL query statement below
SELECT w.id FROM Weather AS w, Weather AS o WHERE w.temperature > o.temperature AND o.recordDate = DATE_SUB(w.recordDate, INTERVAL 1 DAY);