using netponto.AssetManagement.Services;





var builder = WebApplication.CreateBuilder(args);

IConfiguration Configuration = builder.Configuration
     .SetBasePath(Directory.GetCurrentDirectory())
     .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
     .AddJsonFile($"appsettings.{Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT")}.json", optional: true, reloadOnChange: true)
     .AddEnvironmentVariables()
     .Build();

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc();

var app = builder.Build();

app.UseRouting();

// Configure the HTTP request pipeline.
app.UseEndpoints(endpoints =>
    {
        endpoints.MapGrpcService<AssetService>();
        endpoints.MapGrpcService<ProviderService>();

    });
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Run();