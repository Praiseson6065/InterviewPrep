package main


type Video struct{
	Title string
}

func (v *Video) GetTitle() string{
	return v.Title
}

func NewVideo(title string) *Video {
	return &Video{Title: title}
}
type PlaylistIterator interface{
	HasNext() bool
	Next() *Video
}

type Playlist interface{
	createIterator() PlaylistIterator 
}

type YoutubePlaylist struct{
	videos []*Video
}

func (ypl *YoutubePlaylist) AddVideo(video *Video){
	ypl.videos = append(ypl.videos, video)
}

func NewYouTubePlaylist() *YoutubePlaylist {
	return &YoutubePlaylist{
		videos: []*Video{},
	}
}

type YoutubePlaylistIterator struct{
	videos []*Video
	position int
}

func (ypli *YoutubePlaylistIterator) HasNext() bool{
	return ypli.position < len(ypli.videos)
}

func (it *YoutubePlaylistIterator) Next() *Video {
	if !it.HasNext() {
		return nil
	}

	video := it.videos[it.position]
	it.position++
	return video
}


func (p *YoutubePlaylist) CreateIterator() PlaylistIterator {
	return &YoutubePlaylistIterator{
		videos: p.videos,
		position: 0,
	}
}



// func main() {
// 	playlist := NewYouTubePlaylist()

// 	playlist.AddVideo(NewVideo("LLD Tutorial"))
// 	playlist.AddVideo(NewVideo("System Design Basics"))

// 	iterator := playlist.CreateIterator()

// 	for iterator.HasNext() {
// 		fmt.Println(iterator.Next().Title)
// 	}
// }