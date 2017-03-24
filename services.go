package intersights

import (
	"context"
	"os"

	"google.golang.org/grpc/metadata"
)

//GetGrpcLocation - Get the location of a gRPC service
func GetGrpcLocation(vendor string, app string) string {
	location := "http://" + app + "." + vendor + ".intersights.io:50051"
	if addr := os.Getenv("ISGRPC_" + app + "_" + vendor + "_HOST"); addr != "" {
		location = addr
	}
	return location
}

func SetCallContext(ctx context.Context, projectID string) context.Context {
	md, _ := metadata.FromContext(ctx)
	md = md.Copy()
	md["x-intersight-api-client"] = []string{apiClient}
	md["x-intersight-api-version"] = []string{apiVersion}
	md["x-intersight-project"] = []string{projectID}
	return metadata.NewContext(ctx, md)
}
