namespace netponto.AssetManagement;
using Grpc.Core;
    public static class HttpUtils
{
    public static async Task AddHttpHeader(ServerCallContext context, string code) =>
             await context.WriteResponseHeadersAsync(new Metadata()
        {
                    new Metadata.Entry("x-http-code", code)
        });
}

