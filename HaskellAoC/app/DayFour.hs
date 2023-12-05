module DayFour where
import Data.List.Split (splitOn)
import Data.List (intersect)

parseP1 :: String -> [String]
parseP1 = uncurry intersect . break (=="|") . words . drop 9

makeForP2 :: [String] -> [(Int,Int)]
makeForP2 = zipWith (\a b -> (a,length b)) [1..1]

a :: [(Int,Int,Int)] -> Int
a xs = 0
    where helper = 0

solve1 :: String -> Int
solve1 = foldl (\acc x -> if acc == 0 then 1 else acc*2) 0 . parseP1

dayFourP1 :: [String] -> Int
dayFourP1 = foldl (\acc x -> acc + solve1 x) 0