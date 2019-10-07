<?php 

$internal_node = [];
$leaf_node  = [];

// $internal_node = [33, 100];
// $leaf_node = [[10,20,33],[50,70,100],[150,200]];

// 计算内部节点所在索引值
function getInternalPos($num){
    global  $internal_node;

    $postion = 0;
    if(in_array($num, $internal_node)){
        exit('exist num'.PHP_EOL);
    }

    if($internal_node){
        $left = 0;
        $right = count($internal_node) - 1;
        
        while($internal_node[$left] != $internal_node[$right]){
            $mid = intval(($left + $right)/2);
            if($num < $internal_node[$left]){
                return $left;
            }
            if($num > $internal_node[$right]){
                return $right + 1;
            }
            if($num < $internal_node[$mid]){
                $right = $mid -1;
                $postion = $mid;
            }
            if($num > $internal_node[$mid]){
                $left = $mid + 1;
                $postion = $mid + 1;
            }
        }

    }

    return $postion;
}

function insert_node($num)
{
    global  $internal_node;
    global  $leaf_node;

    $internalPretend = $internal_node;
    $leafPretend = $leaf_node;

    $pos = getInternalPos($num);

    $leafArr = $leaf_node[$pos] ?? [];
    if(in_array($num, $leafArr)){
        exit('exist num'.PHP_EOL);
    }

    array_push($leafArr, $num);
    sort($leafArr);
    if(count($leafArr) == 4){
        $leafSplit = array_chunk($leafArr, 2);
        // 重新计算内部节点
        $internalAppend = array_slice($internal_node, $pos + 1);
        array_splice($internalPretend, $pos);

        $internal_node = $internalPretend;
        $internal_node[$pos] = $leafSplit[0][1];
        $internal_node[$pos+1] = $leafSplit[1][1];
        $internal_node = array_merge($internal_node, $internalAppend);

        // 重新计算叶子节点
        $leafAppend = array_slice($leaf_node, $pos + 1);
        array_splice($leafPretend, $pos);

        $leaf_node = $leafPretend;
        $leaf_node[$pos] = $leafSplit[0];
        $leaf_node[$pos+1] = $leafSplit[1];
        $leaf_node = array_merge($leaf_node, $leafAppend);

    }else{
        $leaf_node[$pos] = $leafArr;
    }

}

fwrite(STDOUT, "请输出primary key:".PHP_EOL);

while(1){
    $input = trim(fgets(STDIN));
    if($input == '.btree'){
        fwrite(STDOUT, "btree内部节点:".PHP_EOL);
        print_r($internal_node);
        fwrite(STDOUT, "btree叶子节点:".PHP_EOL);
        print_r($leaf_node);
        continue;
    }

    if(intval($input) <= 0){
        echo "必须输入大于0的整型数字".PHP_EOL;
    }

    insert_node($input);
}
