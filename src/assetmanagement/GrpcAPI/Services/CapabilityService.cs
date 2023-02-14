namespace netponto.AssetManagement.Services;

using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

public class CapabilityService : Protos.Capability.CapabilityBase
{
    public override Task<CreateCapabilityResponse> CreateCapability(CreateCapabilityRequest request, Grpc.Core.ServerCallContext context)
    {
        return base.CreateCapability(request, context);
    }

    public override Task<ListCapabilitiesByTenantResponse> ListCapabilitiesByTenant(ListCapabilitiesByTenantRequest request, Grpc.Core.ServerCallContext context)
    {
        return base.ListCapabilitiesByTenant(request, context); 
    }


}   
