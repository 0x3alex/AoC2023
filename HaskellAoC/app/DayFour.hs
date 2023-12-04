module DayFour where
import Data.List.Split (splitOn)
import Data.List (intersect)

parseP1 :: String -> [String]
parseP1 = uncurry intersect . break (=="|") . words . drop 9

solve1 :: String -> Int
solve1 = foldl (\acc x -> if acc == 0 then 1 else acc*2) 0 . parseP1

dayFourP1 :: [String] -> Int
dayFourP1 = foldl (\acc x -> acc + solve1 x) 0