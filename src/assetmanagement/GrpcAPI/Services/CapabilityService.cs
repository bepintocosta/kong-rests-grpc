namespace netponto.AssetManagement.Services;

using netponto.AssetManagement.Protos;
using System.Threading.Tasks;

public class CapabilityService : Protos.Capability.CapabilityBase
{
    public async override Task<CreateCapabilityResponse> CreateCapability(CreateCapabilityRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "201");
        return new CreateCapabilityResponse() { Url = String.Format("{0}/capability/01", request.TenantCode)};
    }

    public async override Task<ListCapabilitiesResponse> ListCapabilities(ListCapabilitiesRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "200");
        var result = new  ListCapabilitiesResponse() { TotalCount = 1};
        result.Data.Add(new CapabilitieResponse() { Id = 1 , Name = "capability 01" });
        return result;
    }

}   
