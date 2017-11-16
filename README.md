# ecalc
Entropy calculator for binary data (such as PE) files with Go 
 
Example output: 
```D:\Dev\ecalc>ecalc.exe ../ObjectGo.exe
2017/11/16 14:43:15 File data:
2017/11/16 14:43:15     Name: ../ObjectGo.exe
2017/11/16 14:43:15     Size: 1305088
2017/11/16 14:43:15     Entropy: 5.759366365763936
2017/11/16 14:43:15     MD5: c7ae2b72ea285a8f800e7c814bacf5ed
2017/11/16 14:43:15     SHA1: 9e4077d7a55bf9c21cf51e23b1444a5cc82a2d9d
2017/11/16 14:43:15     SHA256: 59e63bc5fb863ddc184a39831097fa8c1bf1fad60d1deeb302557bd360fc7776
2017/11/16 14:43:15 PE header entropy: 1.3454697060501821
2017/11/16 14:43:15 Sections entropy:
2017/11/16 14:43:15     [.text] 5.805650
2017/11/16 14:43:15     [.data] 4.362357
2017/11/16 14:43:15     [.idata] 4.249183
2017/11/16 14:43:15     [.symtab] 0.020393
```

### Known issues 
debug/pe Go package very simple and pretty dumb, so if you try to analyze packed file for example with UPX it will show you smth like `Failed to read PE file header: fail to read string table length: EOF`.  
Need to use (create your own?) different package to analyze PE files (port python's pefile package?).  
