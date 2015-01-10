package classfile

import "fmt"

type ConstantPool struct {
    //cpCount uint
    cpInfos []ConstantInfo
}

func readConstantPool(reader *ClassReader) (*ConstantPool) {
    cpCount := reader.readUint16()
    cpInfos := make([]ConstantInfo, cpCount)

    // The constant_pool table is indexed from 1 to constant_pool_count - 1. 
    for i := uint16(1); i < cpCount; i++ {
        tag := reader.readUint8()
        cpInfos[i] = readConstantInfo(reader, tag)
        // http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
        // All 8-byte constants take up two entries in the constant_pool table of the class file.
        // If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
        // table at index n, then the next usable item in the pool is located at index n+2. 
        // The constant_pool index n+1 must be valid but is considered unusable. 
        if tag == CONSTANT_Long || tag == CONSTANT_Double {
            i++;
        }
    }

    return &ConstantPool{cpInfos}
}

// todo
func (self *ConstantPool) getUtf8(index uint16) (string) {
    cpInfo := self.cpInfos[index]
    if utf8Info, ok := cpInfo.(*ConstantUtf8Info); ok {
        return utf8Info.str
    } 
    
    // todo
    panic(fmt.Sprintf("Const#%v is not ConstantUtf8Info!", index))
}
