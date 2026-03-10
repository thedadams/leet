# Write your MySQL query statement below
SELECT DISTINCT(p.email) FROM Person AS p INNER JOIN Person AS o ON p.email = o.email AND p.id != o.id;