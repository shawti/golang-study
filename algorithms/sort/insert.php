<?php
function InsertSort(array &$s)
{
    for ($i=1; $i < count($s); $i++) {
        $extract = $s[$i];
        for ($j=$i; $j > 0 && $s[$j-1] > $extract; $j--) {
            $tmp = $s[$j-1];
            $s[$j] = $s[$j-1];
            $s[$j-1] = $tmp;
        }
        $s[$j] = $extract;
    }
}

$s = [3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48];
print_r($s);
InsertSort($s);
print_r($s);
