// main.go part 1 of 4
package main
import (
    "fmt"
    "./users"
    flatbuffers "github.com/google/flatbuffers/go"
)

// main.go part 2 of 4
func MakeUser(b *flatbuffers.Builder, name []byte, id uint64) []byte {
    // re-use the already-allocated Builder:
    b.Reset()

    // create the name object and get its offset:
    name_position := b.CreateByteString(name)

    // write the User object:
    users.UserStart(b)
    users.UserAddName(b, name_position)
    users.UserAddId(b, id)
    user_position := users.UserEnd(b)

    // finish the write operations by our User the root object:
    b.Finish(user_position)

    // return the byte slice containing encoded data:
    return b.Bytes[b.Head():]
}

// main.go part 3 of 4
func ReadUser(buf []byte) (name []byte, id uint64) {
    // initialize a User reader from the given buffer:
    user := users.GetRootAsUser(buf, 0)

    // point the name variable to the bytes containing the encoded name:
    name = user.Name()

    // point the id variable to the bytes containing the encoded id:
    id = user.Id()

    return
}

// main.go part 4 of 4
func main() {
    b := flatbuffers.NewBuilder(0)
    buf := MakeUser(b, []byte("Arthur Dent"), 42)
    name, id := ReadUser(buf)
    fmt.Printf("%s has id %d. The encoded data is %d bytes long.\n", name, id, len(buf))
}
