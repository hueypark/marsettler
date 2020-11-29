using UnrealBuildTool;

public class Flatbuffers : ModuleRules
{
	public Flatbuffers(ReadOnlyTargetRules Target) : base(Target)
	{
		PCHUsage = PCHUsageMode.UseExplicitOrSharedPCHs;
	}
}
