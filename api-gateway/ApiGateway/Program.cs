using Ocelot.DependencyInjection;
using Ocelot.Middleware;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Host
  .ConfigureAppConfiguration((hostingContext, config) =>
  {
    config
          .SetBasePath(hostingContext.HostingEnvironment.ContentRootPath)
          .AddJsonFile("appsettings.json", true, true)
          .AddJsonFile($"appsettings.{hostingContext.HostingEnvironment.EnvironmentName}.json", true, true)
          //.AddJsonFile("ocelot.json", false, false)
          .AddJsonFile($"ocelot.{hostingContext.HostingEnvironment.EnvironmentName}.json", true, true)
          .AddEnvironmentVariables();
  })
 .UseDefaultServiceProvider((context, options) =>
{
  options.ValidateOnBuild = false;
  options.ValidateScopes = false;
});

builder.Services.AddOcelot();

builder.Services.AddCors(options =>
{
  options.AddPolicy("CorsPolicy",
      builder => builder.SetIsOriginAllowed((host) => true)
      .AllowAnyMethod()
      .AllowAnyHeader()
      .AllowCredentials());
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
  app.UseDeveloperExceptionPage();
  app.UseSwagger();
  app.UseSwaggerUI();
}

app.MapControllers();

app.UseCors("CorsPolicy");

await app.UseOcelot();

app.Run();
