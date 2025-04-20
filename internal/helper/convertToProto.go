package helper

// import (
// 	"github.com/CelticAlreadyUse/article-story-service/internal/model"
// 	pb "github.com/CelticAlreadyUse/article-story-service/pb/service"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// func ConvertStoriestoProto(data []*model.Story) *pb.Stories {
// 	if data == nil {
// 		return &pb.Stories{}
// 	}
// 	protoStories := make([]*pb.Story, 0, len(data))
// 	for _, story := range data {
// 		if story == nil {
// 			continue
// 		}
// 		protoStories = append(protoStories, convertStoryToProto(story))
// 	}

// 	return &pb.Stories{Stories: protoStories}
// }

// func convertStoryToProto(story *model.Story) *pb.Story {
// 	return &pb.Story{
// 		Id:        story.ID.Hex(),
// 		Title:     story.Title,
// 		AuthorId:  story.AuthorID,
// 		TagsId:    story.TagsID,
// 		Content:   convertContentToProto(story.Content),
// 		CreatedAt: timestamppb.New(story.Created_at),
// 		UpdatedAt: timestamppb.New(story.Updated_at),
// 	}
// }

// func convertContentToProto(contents []*model.StoryElement) []*pb.StoryElement {
// 	if contents == nil {
// 		return nil
// 	}

// 	protoContents := make([]*pb.StoryElement, 0, len(contents))
// 	for _, content := range contents {
// 		if content == nil {
// 			continue
// 		}
// 		protoContents = append(protoContents, convertStoryElementToProto(content))
// 	}

// 	return protoContents
// }

// func convertStoryElementToProto(content *model.StoryElement) *pb.StoryElement {
// 	return &pb.StoryElement{
// 		Type:            string(content.Type),
// 		Text:            content.Text,
// 		Url:             content.ImageUrl,
// 		ImageStyles:     convertImageStylesToProto(content.ImageStyles),
// 		Caption:         content.Caption,
// 		AltText:         content.AltText,
// 		ParagraphStyles: convertParagraphStylesToProto(content.ParagraphStyles),
// 	}
// }

// func convertImageStylesToProto(styles *model.ImageStyles) *pb.ImageStyles {
// 	if styles == nil {
// 		return nil
// 	}
// 	return &pb.ImageStyles{ImageSize: string(styles.ImageSize)}
// }

// func convertParagraphStylesToProto(styles *model.ParagraphStyles) *pb.ParagraphStyles {
// 	if styles == nil {
// 		return nil
// 	}
// 	return &pb.ParagraphStyles{
// 		FontStyle:  string(styles.FontStyle),
// 		FontFamily: string(styles.FontFamily),
// 		FontSize:   string(styles.FontSize),
// 	}
// }
