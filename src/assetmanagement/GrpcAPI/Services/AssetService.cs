using Grpc.Core;
using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

namespace netponto.AssetManagement.Services;

public class AssetService : Protos.Asset.AssetBase
{

    public async override Task<CreateAssetResponse> CreateAsset(CreateAssetRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "201");
        return new CreateAssetResponse() { Url = String.Format("{tenantcode}/assent/01", request.TenantCode)};
    }

    public async override Task<ListAssetsResponse> ListAssets(ListAssetsRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "200");
        var result = new ListAssetsResponse() { TotalCount = 1};
        result.Data.Add(new AssetResponse() { Id = 1 , Name = "provider 01" });
        return result;
    }

   
}
