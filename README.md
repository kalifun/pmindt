# pmindt
> pingcode mind tools

Solve the problem that the mind map is copied to pingcode and the mind map is invalid.

## Install 

```bash
go install github.com/kalifun/pmindt@latest
```

## How to use  

### MindMap to PingCodeMindMap

```bash
pmindt toPCM
```

input mindmap data
```bash
demo
    新建节点1
        node1
            n1
            n2
    新建节点2
        node2
            n3
```

After successful execution, the result will be copied to the clipboard.

```bash
❯ pmindt toPCM      
请输入层次结构（按 Ctrl+D 结束输入）：
demo
    新建节点1
        node1
            n1
            n2
    新建节点2
        node2
            n3
拷贝到PingCode思维导图即可 (已复制到剪贴板）
```

In fact, what we got was a format specified by pingcode.

```bash
<meta charset='utf-8'><html><head></head><body><plait>{"type":"elements","data":[{"id":"eolvck6r1dz4","data":{"topic":{"children":[{"text":"demo"}]}},"children":[{"id":"eolvck6r1dzc","data":{"topic":{"children":[{"text":"新建节点1"}]}},"children":[{"id":"eolvck6r1dzb","data":{"topic":{"children":[{"text":"node1"}]}},"children":[{"id":"eolvck6r1dz8","data":{"topic":{"children":[{"text":"n1"}]}},"children":[],"width":16,"height":20},{"id":"eolvck6r1dza","data":{"topic":{"children":[{"text":"n2"}]}},"children":[],"width":16,"height":20}],"width":40,"height":20}],"width":104,"height":20},{"id":"eolvck6r1dzi","data":{"topic":{"children":[{"text":"新建节点2"}]}},"children":[{"id":"eolvck6r1dzh","data":{"topic":{"children":[{"text":"node2"}]}},"children":[{"id":"eolvck6r1dzg","data":{"topic":{"children":[{"text":"n3"}]}},"children":[],"width":16,"height":20}],"width":40,"height":20}],"width":104,"height":20}],"width":40,"height":25,"layout":"right","rightNodeCount":0,"isRoot":true,"type":"mindmap","points":[[0,12]]}]}</plait></body></html>
```