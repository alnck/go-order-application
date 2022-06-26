#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build

WORKDIR /api-gateway

COPY /api-gateway/ApiGateway/*.csproj .
RUN dotnet restore "ApiGateway.csproj"

COPY /api-gateway/ApiGateway .

WORKDIR "/api-gateway"
RUN dotnet build "ApiGateway.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "ApiGateway.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app

EXPOSE 5000
ENV ASPNETCORE_URLS=http://+:5000

COPY --from=publish /app/publish .

ENTRYPOINT ["dotnet", "ApiGateway.dll"]
