package grpchandler

// import (
// 	"context"
// 	"fmt"

// 	"github.com/CelticAlreadyUse/article-story-service/internal/helper"
// 	"github.com/CelticAlreadyUse/article-story-service/internal/model"
// 	pb "github.com/CelticAlreadyUse/article-story-service/pb/service"
// )

// type StorygRPCHandler struct {
// 	pb.UnimplementedStoryServiceServer
// 	storyUsecase model.StoryUsecase
// }

// func InitStoryGrpcHandler(usecase model.StoryUsecase) pb.StoryServiceServer {
// 	return &StorygRPCHandler{storyUsecase: usecase}
// }

// func (h *StorygRPCHandler) FindAllStoriesByUserID(ctx context.Context, req *pb.GetStoriesByUserIDRequest) (*pb.StoriesResponse,error) {
// 	if req.Id < 0 {
// 		return nil, fmt.Errorf("invalid user id")
// 	}

// 	Stories,cursor,err := h.storyUsecase.GetStoriesByUserID(ctx, int64(req.Id),req.Cursor)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get stories")
// 	}
	
// 	storiesProto := helper.ConvertStoriestoProto(Stories)
// 	return  storiesProto,nil
// }


