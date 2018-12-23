# introduce
It is a music player. It implements parts of features, which contains(add, delete, list, play). It is a simple player.

# How to build
+ 1, create a project: mkdir /projectDir/player;
+ 2, copy main into this dir(cd /projectDir/player; cp main.go ./)
+ 3, create a src direcory: mdkir /srcDir/music
+ 4, copy source file into this dir(cd /srcDir/music; cp music*.go ./)

# How to test
+ After building, test the main program. Rebuild the project(main), and test. (go build main.go; ./main)
+ After/before building, make the unit test. Go to the src directory. (cd /srcDir/music; go test)

# Feature introduce
+ It only support simple features:
+ Add music;
+ Play music;
+ Remove music;
+ List all music;

# Output sample
***@localhost:~/go/player$ ./main
L(list music); P <index>(playing); A <singer>:<song name>; D <index>(deleting music); X(exit the program)
Please input:
A Beyonce Halo
Please input:
A Muziya Possible   
Please input:
L
index: 0 song-name: Halo
index: 1 song-name: Possible
Please input:
P 0
[P 0]
It is playing:(Halo)<Beyonce>
Please input:
P 1
[P 1]
It is playing:(Possible)<Muziya>
Please input:
L
index: 0 song-name: Halo
index: 1 song-name: Possible
Please input:
D 0
[{Beyonce Halo} {Muziya Possible}]
[{Muziya Possible}]
Please input:
L     
index: 0 song-name: Possible
Please input:
X
It is over, and program exits!
