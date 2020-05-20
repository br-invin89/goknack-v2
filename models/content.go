package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Tag struct
type Tag struct {
	ID						bson.ObjectId		`json:"id" bson:"_id"`
	PrimaryMarket			string				`json:"primaryMarket" bson:"primaryMarket"`
	Description				string				`json:"description" bson:"description"`
	IsActive				bool				`json:"isActive" bson:"isActive"`
	AddedBy					string				`json:"addedBy" bson:"addedBy"`
	AddedByID				string				`json:"addedByID" bson:"addedByID"`
	DateAdded				time.Time			`json:"dateAdded" bson:"dateAdded"`
}

// Content struct
type Content struct {
	ID						bson.ObjectId		`json:"id" bson:"_id"`
	FilePath				string				`json:"filePath" bson:"filePath"`
	ThumbPath				string				`json:"thumbPath" bson:"thumbPath"`
	FileName				string				`json:"fileName" bson:"fileName"`
	FileSize				int					`json:"fileSize" bson:"fileSize"`
	FileType				string				`json:"fileType" bson:"fileType"`
	IsImage					bool				`json:"isImage" bson:"isImage"`
	IsAudio					bool				`json:"isAudio" bson:"isAudio"`
	Title					string				`json:"title" bson:"title"`
	Description				string				`json:"description" bson:"description"`
	Tags					[]Tag				`json:"tags" bson:"tags"`
	IsEligible				bool				`json:"isEligible" bson:"isEligible"`
	IsActive				bool				`json:"isActive" bson:"isActive"`
	IsDeleted				bool				`json:"isDeleted" bson:"isDeleted"`
	IsFlagged				bool				`json:"isFlagged" bson:"isFlagged"`
	IsPaidOut				bool				`json:"isPaidOut" bson:"isPaidOut"`
	Owner					string				`json:"owner" bson:"owner"`
	OwnerName				string				`json:"ownerName" bson:"ownerName"`
	OwnerEmail				string				`json:"ownerEmail" bson:"ownerEmail"`
	DateUpload				time.Time			`json:"dateUpload" bson:"dateUpload"`
	DateDelete				time.Time			`json:"dateDelete" bson:"dateDelete"`
	DateFlagged				time.Time			`json:"dateFlagged" bson:"dateFlagged"`
	DatePaid				time.Time			`json:"datePaid" bson:"datePaid"`
	DateEligible			time.Time			`json:"dateEligible" bson:"dateEligible"`
}

// SaveContent func
func SaveContent(c Content) error{

	mongo := Session.Copy()
	defer mongo.Close()

	err := mongo.DB("goknack").C("content").Insert(c)

	if err != nil {
		return err
	}

	return nil

}



