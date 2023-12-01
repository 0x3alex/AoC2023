module DayOne (dayOneP1, dayOneP2) where

import Data.Char (isDigit, digitToInt, intToDigit)
import qualified Data.Map as Map
import Data.Maybe (isNothing, fromJust, fromMaybe)
import Utils.Containers.Internal.StrictPair (StrictPair)
import Data.List (isInfixOf)


stringToInt :: String -> Int
stringToInt = read

numNames :: [String]
numNames = ["one","two","three","four","five","six","seven","eight","nine"]

lookUp :: String -> Char
lookUp s = fromMaybe ' ' $ Map.lookup s $ Map.fromList $ zip numNames [intToDigit x | x <- [1..]]

f :: String -> String -> String
f [] _ = []
f k@(x:xs) a
    | isDigit x = x : f xs ""
    | not . null . m $ (a++[x]) = lookUp (head (m (a++[x]))) : f k ""
    | otherwise = f xs (a++[x])
    where   m ks = filter (`isInfixOf` ks) numNames

dayOneP2 :: [String] -> Int
dayOneP2 = dayOneP1 . map (`f` "") 

dayOneP1 :: [String] -> Int
dayOneP1 = foldr (\ x -> (+) (buildDigits . extractDigits $ x)) 0
    where
        extractDigits = filter isDigit
        buildDigits s
            | length s == 1 = stringToInt (s ++ s)
            | length s > 2 = stringToInt (head s : [last s])
            | otherwise = stringToInt s
