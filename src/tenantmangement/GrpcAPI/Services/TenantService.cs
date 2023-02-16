using Grpc.Core;
using netponto.TenantManagement.Protos;
using System.Threading.Tasks;

namespace netponto.TenantManagement.Services;

public class TenantService : Protos.Tenant.TenantBase
{
    public async override Task<CreateTenantResponse> CreateTenant(CreateTenantRequest request, Grpc.Core.ServerCallContext context)
    {
        await HttpUtils.AddHttpHeader(context, "201");
        return new CreateTenantResponse(){ Url = "/tenant/test"};
    }
}
