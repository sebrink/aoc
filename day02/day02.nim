import strutils, sequtils, sugar, os, sets, tables

proc partOne(file: seq[string]): int =
    var restab = initTable[int, int]()
    var checksum = 0
    for line in file:
        var iTable = initTable[char,int]()
        for i in line:
            if iTable.hasKey(i):
                iTable[i] += 1
            else:
                iTable[i] = 1
        var interSet = initSet[int]()
        for x, y in iTable:
            interSet.incl y
        for y in interset:
            if y != 1:
                if restab.hasKey(y):
                    restab[y] += 1
                else:
                    restab[y] = 1
    result = 1
    for x, y in restab:
        result *= y

let file = readFile("input.txt").splitLines.filterIt(it.len > 0)
echo file.partOne