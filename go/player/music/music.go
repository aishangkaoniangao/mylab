package music

import "fmt"

type song struct {
    singer string
    name string
}

type SongLib struct {
    songs []song
}

func NewLib() (*SongLib) {
    sl := SongLib{}
    return &sl
}

func (self *SongLib) Delete(index int) {
    l := len(self.songs)
    if (index < l && index >= 0) {
        if (index == 0) {
            self.songs = self.songs[1:]
            return
        } else if (l == index + 1) {
            self.songs = self.songs[:index] //go的切片[a:b] a是开始下标 b是结束下标 但是包含a不包含b
        } else {
            self.songs = append(self.songs[:index-1], self.songs[index + 1])
        }
    }
}

func (self *SongLib) Len() (int){
    return len(self.songs)
}

func (self *SongLib) Show() {
    for i :=0; i <len(self.songs); i++ {
        fmt.Println("index:", i, "song-name:", self.songs[i].name)
    }
}

func (self *SongLib) Add(singer string, name string) {
    s := song{}
    s.singer = singer
    s.name = name
    self.songs = append(self.songs, s)
}

func (self *SongLib) Play(index int) {
    if (index >=0 && index < len(self.songs)) {
        fmt.Println(fmt.Sprintf("It is playing:(%s)<%s>", self.songs[index].name, self.songs[index].singer))
    }
}