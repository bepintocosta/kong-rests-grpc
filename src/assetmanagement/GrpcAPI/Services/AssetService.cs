using Grpc.Core;
using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

namespace netponto.AssetManagement.Services;

public class AssetService : Protos.Asset.AssetBase
{


    public async override Task<ListAssetsResponse> ListAssets(ListAssetsRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "200");
        return new ListAssetsResponse();
    }

    public override Task<CreateAssetResponse> CreateAsset(CreateAssetRequest request, Grpc.Core.ServerCallContext context)
    {
        return base.CreateAsset(request, context);
    }
}
