package grpchandler

import (
	"context"
	"fmt"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	pb "github.com/CelticAlreadyUse/article-story-service/pb/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StorygRPCHandler struct {
	pb.UnimplementedStoryServiceServer
	storyUsecase model.StoryUsecase
}

func InitStoryGrpcHandler(usecase model.StoryUsecase) pb.StoryServiceServer {
	return &StorygRPCHandler{storyUsecase: usecase}
}

func (h *StorygRPCHandler) FindAllStoriesByUserID(ctx context.Context, req *pb.GetStoriesByUserIDRequest) (*pb.Stories, error) {
	if req.Id < 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	Stories, err := h.storyUsecase.GetStoriesByUserID(ctx, int64(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to get stories")
	}
	storiesProto := convertStoriestoProto(Stories)
	return  storiesProto,nil
}


func convertStoriestoProto(data []*model.Story) *pb.Stories {
	if data == nil {
		fmt.Println("convertStoriestoProto: received nil data")
		return &pb.Stories{Stories: []*pb.Story{}}
	}

	var protoStories []*pb.Story

	for _, story := range data {
		if story == nil {
			fmt.Println("convertStoriestoProto: found nil story, skipping")
			continue
		}

		var protoContent []*pb.StoryElement
		for _, content := range story.Content {
			if content == nil {
				fmt.Println("convertStoriestoProto: found nil content, skipping")
				continue
			}

			// Default value untuk ImageStyles
			var imageStyles *pb.ImageStyles
			if content.ImageStyles != nil {
				imageStyles = &pb.ImageStyles{ImageSize: string(content.ImageStyles.ImageSize)}
			}

			// Default value untuk ParagraphStyles
			var paragraphStyles *pb.ParagraphStyles
			if content.ParagraphStyles != nil {
				paragraphStyles = &pb.ParagraphStyles{
					FontStyle:  string(content.ParagraphStyles.FontSize),
					FontFamily: string(content.ParagraphStyles.FontFamily),
					FontSize:   string(content.ParagraphStyles.FontSize),
				}
			}

			protoContent = append(protoContent, &pb.StoryElement{
				Type:            string(content.Type),
				Text:            content.Text,
				Url:             content.Url,
				ImageStyles:     imageStyles,
				Caption:         content.Caption,
				AltText:         content.AltText,
				ParagraphStyles: paragraphStyles,
			})
		}

		protoStories = append(protoStories, &pb.Story{
			Id:        story.ID.Hex(),
			Title:     story.Title,
			AuthorId:  story.AuthorID,
			Tags:      story.Tags,
			Content:   protoContent,
			CreatedAt: timestamppb.New(story.Created_at),
			UpdatedAt: timestamppb.New(story.Updated_at),
		})
	}

	return &pb.Stories{Stories: protoStories}
}

