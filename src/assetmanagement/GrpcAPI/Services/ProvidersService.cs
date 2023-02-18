namespace netponto.AssetManagement.Services;

using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

public class ProviderService : Protos.Provider.ProviderBase
{

    public async override Task<CreateProviderResponse> CreateProvider(CreateProviderRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "201");
        return new CreateProviderResponse() { Url = String.Format("{0}/provider/01", request.TenantCode)};
    }

    public async override Task<ListProvidersResponse> ListProviders(ListProvidersRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "200");
        var result = new ListProvidersResponse(){TotalCount = 1 };
        result.Data.Add( new ProviderResponse {Id = 1 , Name = "provider 01"} );
        return result;
    }

    
}
