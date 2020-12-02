// Copyright Epic Games, Inc. All Rights Reserved.

using UnrealBuildTool;

public class Marsettler : ModuleRules
{
	public Marsettler(ReadOnlyTargetRules Target) : base(Target)
	{
		PCHUsage = PCHUsageMode.UseExplicitOrSharedPCHs;

		PublicDependencyModuleNames.AddRange(new string[] {
			"AIModule",
			"Core",
			"CoreUObject",
			"Engine",
			"HeadMountedDisplay",
			"InputCore",
			"NavigationSystem",
			"Networking",
			"Sockets" });

		PublicIncludePaths.Add("Marsettler");
	}
}
