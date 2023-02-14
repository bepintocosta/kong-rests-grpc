namespace netponto.AssetManagement.Services;

using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

public class ProviderService : Protos.Provider.ProviderBase
{
    public override Task<ListProvidersResponse> ListProviders(ListProvidersRequest request, Grpc.Core.ServerCallContext context)
    {
        return base.ListProviders(request, context);
    }

    public override Task<CreateProviderResponse> CreateProvider(CreateProviderRequest request, Grpc.Core.ServerCallContext context)
    {
        return base.CreateProvider(request, context);
    }
}
