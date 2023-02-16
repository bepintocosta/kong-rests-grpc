#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app

EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build

COPY ["/src/tenantmangement/GrpcAPI/", "src/tenantmangement/GrpcApi/"]

RUN dotnet restore "/src/tenantmangement/GrpcApi/GrpcApi.csproj"

COPY . .

WORKDIR "src/tenantmangement/GrpcApi"

RUN dotnet build "GrpcApi.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "GrpcApi.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "GrpcApi.dll"]