-- MOHAMMAD RAZI UL HAQ
-- 2020-10-0015

-- README
-- Part 1, 2, 5, 6 work. Didn't attemp 3,4 :(. The ones I did work
-- 	according to what i could understand from the comments and
-- 	writeup. Hope what i understood was okay. Happy Checking :))
-- ---------------------------------------------------------------------
-- DNA Analysis 
-- CS300 Spring 2018
-- Due: 24 Feb 2018 @9pm
-- ------------------------------------Assignment 2------------------------------------
--
-- >>> YOU ARE NOT ALLOWED TO IMPORT ANY LIBRARY
-- Functions available without import are okay
-- Making new helper functions is okay
--
-- ---------------------------------------------------------------------
--
-- DNA can be thought of as a sequence of nucleotides. Each nucleotide is 
-- adenine, cytosine, guanine, or thymine. These are abbreviated as A, C, 
-- G, and T.
--
type DNA = [Char]
type RNA = [Char]
type Codon = [Char]
type AminoAcid = Maybe String

-- ------------------------------------------------------------------------
-- 				PART 1
-- ------------------------------------------------------------------------				

-- We want to calculate how alike are two DNA strands. We will try to 
-- align the DNA strands. An aligned nucleotide gets 3 points, a misaligned
-- gets 2 points, and inserting a gap in one of the strands gets 1 point. 
-- Since we are not sure how the next two characters should be aligned, we
-- try three approaches and pick the one that gets us the maximum score.
-- 1) Align or misalign the next nucleotide from both strands
-- 2) Align the next nucleotide from first strand with a gap in the second     
-- 3) Align the next nucleotide from second strand with a gap in the first    
-- In all three cases, we calculate score of leftover strands from recursive 
-- call and add the appropriate penalty.                                    

val :: DNA -> DNA -> Int
val [] [] = 0

val (x:xs) []      	|(x/='-')		= 1 + val xs []
val (x:xs) []      	| otherwise 	= 0
val [] 	 (y:ys)  	|(y/='-')		= 1 + val [] ys
val [] 	 (y:ys)  	| otherwise 	= 0

val (x:xs) (y:ys) 
	| (x==y),(x/='-'), (y/='-') 	= 3 + maximum ([val xs ys] ++ [val xs ('-':ys)] ++ [val ('-':xs) ys]) 
	| (x=='-'), (y/='-') 			= 1 + maximum ([val xs ys] ++ [val xs ('-':ys)] ++ [val ('-':xs) ys]) 
	| (y=='-'), (x/='-') 			= 1 + maximum ([val xs ys] ++ [val xs ('-':ys)] ++ [val ('-':xs) ys]) 
	| (x/=y) 						= 2 + maximum ([val xs ys] ++ [val xs ('-':ys)] ++ [val ('-':xs) ys]) 

score :: DNA -> DNA -> Int
score (x:xs) (y:ys) = maximum ([val (x:xs) (y:ys)] ++ [val ('-':x:xs) (y:ys)] ++ [val (x:xs) ('-':y:ys)])


-- ['A','T','C','C'] ['A','C','C','G']

-- Add condition that cannot insert 2 spaces together

-- -------------------------------------------------------------------------
--				PART 2
-- -------------------------------------------------------------------------
-- Write a function that takes a list of DNA strands and returns a DNA tree. 
-- For each DNA strand, make a separate node with zero score 
-- in the Int field. Then keep merging the trees. The merging policy is:
-- 	1) Merge two trees with highest score. Merging is done by making new
--	node with the smaller DNA (min), and the two trees as subtrees of this
--	tree
--	2) Goto step 1 :)
--


data DNATree = Node DNA Int DNATree DNATree | Nil deriving (Ord, Show, Eq)

makeDNATree :: [DNA] -> DNATree
makeDNATree (x:xs) = head(applyNTimes ((length (x:xs)) -1) final values)
	where values = makeTreeList (x:xs)


z =(makeTreeList (["AAAAAAAA", "AAAAAAAT", "AAAAAATT", "AAAAATTT"]))
f =  callNscore z

-- On each call adds the max score node and removes its children from the list
final :: [DNATree] -> [DNATree]
final list =([maxi(callNscore(list))] ++ (toRemove (maxi(callNscore(list))) list))

-- Removes the trees of the newly formed node from the list
toRemove (Node dnax scorex (Node dnal scorel leftl rightl) (Node dnar scorer leftr rightr)) list = (remv dnal (remv dnar list))

-- Returns the max score node from list of all possible node configurations
maxi :: [DNATree] ->DNATree
maxi [x] = x
maxi ((Node dnax scorex leftx rightx):(Node dnay scorey lefty righty):xs)
	| scorex>=scorey = maxi ((Node dnax scorex leftx rightx) : xs)
	| otherwise 	 = maxi ((Node dnay scorey lefty righty):xs)


--Removes those trees from the list which are already children of a tree in the list
remv _ [] = []
remv x ((Node dnay scorey lefty righty):xs)
	| x==dnay = xs
	|otherwise = [(Node dnay scorey lefty righty)] ++ remv x xs


-- Gives all possible combination of trees that could be made
callNscore:: [DNATree] -> [DNATree]
callNscore [x] = []
callNscore (x:xs) = (scoreNcalc (x:xs) ++ callNscore (xs))

scoreNcalc:: [DNATree] -> [DNATree]
scoreNcalc [] = []
scoreNcalc [a] = []
scoreNcalc ((Node dnax scorex leftx rightx):(Node dnay scorey lefty righty):xs) = [mergeTwo (Node dnax scorex leftx rightx) (Node dnay scorey lefty righty)] ++ scoreNcalc ((Node dnax scorex leftx rightx):xs)


-- Forms the initial 0 score tree list from [DNA]
makeTreeList :: [DNA] -> [DNATree]
makeTreeList [] 	= []
makeTreeList (x:xs) = [Node x 0 Nil Nil] ++ makeTreeList xs

mergeTwo :: DNATree -> DNATree -> DNATree
mergeTwo (Node dnax scorex leftx rightx) (Node dnay scorey lefty righty) 
	| min dnax dnay == dnax = Node (min dnax dnay) (score dnax dnay) (Node dnax scorex leftx rightx) (Node dnay scorey lefty righty)
	| otherwise = Node (min dnax dnay) (score dnax dnay) (Node dnay scorey lefty righty) (Node dnax scorex leftx rightx) 

applyNTimes :: Int -> (a -> a) -> a -> a
applyNTimes n f x 
    | n == 0        = x
    | otherwise     = f (applyNTimes (n-1) f x)


-- -------------------------------------------------------------------------
--				PART 3
-- -------------------------------------------------------------------------

-- Even you would have realized it is hard to debug and figure out the tree
-- in the form in which it currently is displayed. Lets try to neatly print 
-- the DNATree. Each internal node should show the 
-- match score while leaves should show the DNA strand. In case the DNA strand 
-- is more than 10 characters, show only the first seven followed by "..." 
-- The tree should show like this for an evolution tree of
-- ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]
--
-- 20
-- +---ATATTATA
-- +---21
--     +---21
--     |   +---ACTGCATG
--     |   +---ACTACACC
--     +---AACCTTGG
--
-- Make helper functions as needed. It is a bit tricky to get it right. One
-- hint is to pass two extra string, one showing what to prepend to next 
-- level e.g. "+---" and another to prepend to level further deep e.g. "|   "

-- draw :: DNATree -> [Char]
-- draw = undefined

-- ---------------------------------------------------------------------------
--				PART 4
-- ---------------------------------------------------------------------------
--
--
-- Our score function is inefficient due to repeated calls for the same 
-- suffixes. Lets make a dictionary to remember previous results. First you
-- will consider the dictionary as a list of tuples and write a lookup
-- function. Return Nothing if the element is not found. Also write the 
-- insert function. You can assume that the key is not already there.
-- type Dict a b = [(a,b)]

-- lookupDict :: (Eq a) => a -> Dict a b -> Maybe b
-- lookupDict = undefined

-- insertDict :: (Eq a) => a -> b -> (Dict a b)-> (Dict a b)
-- insertDict = undefined

-- We will improve the score function to also return the alignment along
-- with the score. The aligned DNA strands will have gaps inserted. You
-- can represent a gap with "-". You will need multiple let expressions 
-- to destructure the tuples returned by recursive calls.

-- alignment :: String -> String -> ((String, String), Int)
-- alignment = undefined

-- We will now pass a dictionary to remember previously calculated scores 
-- and return the updated dictionary along with the result. Use let 
-- expressions like the last part and pass the dictionary from each call
-- to the next. Also write logic to skip the entire calculation if the 
-- score is found in the dictionary. You need just one call to insert.
-- type ScoreDict = Dict (DNA,DNA) Int

-- scoreMemo :: (DNA,DNA) -> ScoreDict -> (ScoreDict,Int)
-- scoreMemo = undefined
-- In this part, we will use an alternate representation for the 
-- dictionary and rewrite the scoreMemo function using this new format.
-- The dictionary will be just the lookup function so the dictionary 
-- can be invoked as a function to lookup an element. To insert an
-- element you return a new function that checks for the inserted
-- element and returns the old dictionary otherwise. You will have to
-- think a bit on how this will work. An empty dictionary in this  
-- format is (\_->Nothing)

-- type Dict2 a b = a->Maybe b
-- insertDict2 :: (Eq a) => a -> b -> (Dict2 a b)-> (Dict2 a b)
-- insertDict2 = undefined

-- type ScoreDict2 = Dict2 (DNA,DNA) Int

-- scoreMemo2 :: (DNA,DNA) -> ScoreDict2 -> (ScoreDict2,Int)
-- scoreMemo2 = undefined

-- ---------------------------------------------------------------------------
-- 				PART 5
-- ---------------------------------------------------------------------------

-- Now, we will try to find the mutationDistance between two DNA sequences.
-- You have to calculate the number of mutations it takes to convert one 
-- (start sequence) to (end sequence). You will also be given a bank of 
-- sequences. However, there are a couple of constraints, these are as follows:

-- 1) The DNA sequences are of length 8
-- 2) For a sequence to be a part of the mutation distance, it must contain 
-- "all but one" of the neuclotide bases as its preceding sequence in the same 
-- order AND be present in the bank of valid sequences
-- 'AATTGGCC' -> 'AATTGGCA' is valid only if 'AATTGGCA' is present in the bank
-- 3) Assume that the bank will contain valid sequences and the start sequence
-- may or may not be a part of the bank.
-- 4) Return -1 if a mutation is not possible

	
-- mutationDistance "AATTGGCC" "TTTTGGCA" ["AATTGGAC", "TTTTGGCA", "AAATGGCC", "TATTGGCC", "TTTTGGCC"] == 3
-- mutationDistance "AAAAAAAA" "AAAAAATT" ["AAAAAAAA", "AAAAAAAT", "AAAAAATT", "AAAAATTT"] == 2

-- I also make sure to send the new list each time which doesn't have
-- the previous value from which it converted from so that there is
-- no infinite call etc.

mutationDistance :: DNA -> DNA -> [DNA] -> Int
mutationDistance initial final list = checkEach (listChanges initial list) final (filter (initial/=) list) 0

-- This function gets a list of DNA's to check. It sends to checkEqual each one of those.
-- And outputs the max value of checkEqual. Since it returns -1 when there are no more 
-- elements that can be made from the DNA LIST by changing 1 letter in the DNA  hence
-- the max value will be the only positive value in the list. 
checkEach [] final list num = -1
checkEach (x:xs) final list num = maximum([checkEqual x final list num] ++ [checkEach xs final list num])


-- This one gets 1 element from the list of DNA and it first
-- checks whether its the end. Or it sends into recursive call
-- a nother list of DNA's that are different from this by 1.
checkEqual x final list num
	| x==final  = (num + 1)
	| otherwise = checkEach (listChanges x list) final (filter (x/=) list) (num + 1)


-- Gives the number of how different 2 DNA's are
changes :: DNA -> DNA -> Int ->Int
changes [] [] count = count
changes (x:xs) (y:ys) count 
	| (x==y) 	= changes xs ys count
	| (x/=y)   	= changes xs ys (count+1)


-- Gives the list of all dna's that are only 1 different from the initial
listChanges :: DNA -> [DNA] -> [DNA]
listChanges ini [] = []
listChanges ini (l:ls)
		| (changes ini l 0) == 1		= [l] ++ listChanges ini ls
		| otherwise					= listChanges ini ls


-- ---------------------------------------------------------------------------
-- 				PART 6
-- ---------------------------------------------------------------------------
--
-- Now, we will write a function to transcribe DNA to RNA. 
-- The difference between DNA and RNA is of just one base i.e.
-- instead of Thymine it contains Uracil. (U)
--
transcribeDNA :: DNA -> RNA
transcribeDNA [] = []
transcribeDNA  (x:xs)
	| x=='T' 	= ['U'] ++ transcribeDNA xs
	| otherwise = [x] ++ transcribeDNA xs

-- Next, we will translate RNA into proteins. A codon is a group of 3 neuclotides 
-- and forms an aminoacid. A protein is made up of various amino acids bonded 
-- together. Translation starts at a START codon and ends at a STOP codon. The most
-- common start codon is AUG and the three STOP codons are UAA, UAG and UGA.
-- makeAminoAcid should return Nothing in case of a STOP codon.
-- Your translateRNA function should return a list of proteins present in the input
-- sequence. 
-- Please note that the return type of translateRNA is [String], you should convert
-- the abstract type into a concrete one.
-- You might wanna use the RNA codon table from 
-- https://www.news-medical.net/life-sciences/RNA-Codons-and-DNA-Codons.aspx
-- ("UGA", Nothing),("UAA", Nothing),("UAG", Nothing)

dict = [("AUG",Just "Met"),("UUA", Just "Leu"),("UGG", Just "Trp"),("CCU", Just "Pro"),("UCG", Just "Ser"),("UGA", Nothing),("UAA", Nothing),("UAG", Nothing)]
makeAminoAcid :: Codon -> AminoAcid
makeAminoAcid x = search' x dict

search' _ [] = Nothing
search' x ((a,Just b):xs) = if x == a then Just b else search' x xs
search' x ((a,Nothing):xs) = if x == a then Nothing else search' x xs


translateRNA :: RNA -> [String] 
translateRNA [] = []
translateRNA (x1:x2:x3:xs) 
	| (x1:x2:[x3])=="AUG" = ["Met"]++ start (xs)
	| otherwise = translateRNA (xs)

start :: RNA -> [String]
start [] = []
start (x1:x2:x3:xs) 
	| (x1:x2:[x3])=="UGA"= []
	| (x1:x2:[x3])=="UAA"= []
	| (x1:x2:[x3])=="UAG"= []
	| otherwise 		 = [converter (makeAminoAcid (x1:x2:[x3]))] ++ start(xs)



converter :: Maybe String -> [Char]
converter (Just a) = a