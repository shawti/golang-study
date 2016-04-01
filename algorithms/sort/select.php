<?php
function selectSort(array &$s)
{
    $len = count($s);
    $sortCount = 0;
    for ($i=0; $i < $len-1; $i++) {
        $currentMinIndex = $i;
        for ($j=$sortCount; $j < $len; $j++) {
            if($s[$j] < $s[$currentMinIndex]) {
                $currentMinIndex = $j;
            }
        }
        echo 'Minimum: '.$s[$currentMinIndex].PHP_EOL;
        $tmp = $s[$i];
        $s[$i] = $s[$currentMinIndex];
        $s[$currentMinIndex] = $tmp;
        $sortCount++;
        print_r($s);
    }
}

$s = [3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48];
print_r($s);
selectSort($s);
print_r($s);
