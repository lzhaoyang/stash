package models

type Repository interface {
	Gallery() GalleryReaderWriter
	Image() ImageReaderWriter
	Movie() MovieReaderWriter
	Performer() PerformerReaderWriter
	Scene() SceneReaderWriter
	SceneMarker() SceneMarkerReaderWriter
	ScrapedItem() ScrapedItemReaderWriter
	Studio() StudioReaderWriter
	Tag() TagReaderWriter
}

type ReaderRepository interface {
	Gallery() GalleryReader
	Image() ImageReader
	Movie() MovieReader
	Performer() PerformerReader
	Scene() SceneReader
	SceneMarker() SceneMarkerReader
	ScrapedItem() ScrapedItemReader
	Studio() StudioReader
	Tag() TagReader
}
