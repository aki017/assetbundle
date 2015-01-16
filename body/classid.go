package body

type ClassID uint32

const (
	Null                     ClassID = 0
	GameObject                       = 1
	Component                        = 2
	LevelGameManager                 = 3
	Transform                        = 4
	TimeManager                      = 5
	GlobalGameManager                = 6
	Behaviour                        = 8
	GameManager                      = 9
	AudioManager                     = 11
	ParticleAnimator                 = 12
	InputManager                     = 13
	EllipsoidParticleEmitter         = 15
	Pipeline                         = 17
	EditorExtension                  = 18
	Camera                           = 20
	Material                         = 21
	MeshRenderer                     = 23
	Renderer                         = 25
	ParticleRenderer                 = 26
	Texture                          = 27
	Texture2D                        = 28
	Scene                            = 29
	RenderManager                    = 30
	MeshFilter                       = 33
	OcclusionPortal                  = 41
	Mesh                             = 43
	Skybox                           = 45
	QualitySettings                  = 47
	Shader                           = 48
	TextAsset                        = 49
	NotificationManager              = 52
	Rigidbody                        = 54
	PhysicsManager                   = 55
	Collider                         = 56
	Joint                            = 57
	HingeJoint                       = 59
	MeshCollider                     = 64
	BoxCollider                      = 65
	AnimationManager                 = 71
	AnimationClip                    = 74
	ConstantForce                    = 75
	WorldParticleCollider            = 76
	TagManager                       = 78
	AudioListener                    = 81
	AudioSource                      = 82
	AudioClip                        = 83
	RenderTexture                    = 84
	MeshParticleEmitter              = 87
	ParticleEmitter                  = 88
	Cubemap                          = 89
	GUILayer                         = 92
	ScriptMapper                     = 94
	TrailRenderer                    = 96
	DelayedCallManager               = 98
	TextMesh                         = 102
	RenderSettings                   = 104
	Light                            = 108
	CGProgram                        = 109
	Animation                        = 111
	MonoBehaviour                    = 114
	MonoScript                       = 115
	MonoManager                      = 116
	Texture3D                        = 117
	Projector                        = 119
	LineRenderer                     = 120
	Flare                            = 121
	Halo                             = 122
	LensFlare                        = 123
	FlareLayer                       = 124
	HaloLayer                        = 125
	NavMeshLayers                    = 126
	HaloManager                      = 127
	Font                             = 128
	PlayerSettings                   = 129
	NamedObject                      = 130
	GUITexture                       = 131
	GUIText                          = 132
	GUIElement                       = 133
	PhysicMaterial                   = 134
	SphereCollider                   = 135
	CapsuleCollider                  = 136
	SkinnedMeshRenderer              = 137
	FixedJoint                       = 138
	RaycastCollider                  = 140
	BuildSettings                    = 141
	AssetBundle                      = 142
	CharacterController              = 143
	CharacterJoint                   = 144
	SpringJoint                      = 145
	WheelCollider                    = 146
	ResourceManager                  = 147
	NetworkView                      = 148
	NetworkManager                   = 149
	PreloadData                      = 150
	MovieTexture                     = 152
	ConfigurableJoint                = 153
	TerrainCollider                  = 154
	MasterServerInterface            = 155
	TerrainData                      = 156
	LightmapSettings                 = 157
	WebCamTexture                    = 158
	EditorSettings                   = 159
	InteractiveCloth                 = 160
	ClothRenderer                    = 161
	SkinnedCloth                     = 163
	AudioReverbFilter                = 164
	AudioHighPassFilter              = 165
	AudioChorusFilter                = 166
	AudioReverbZone                  = 167
	AudioEchoFilter                  = 168
	AudioLowPassFilter               = 169
	AudioDistortionFilter            = 170
	AudioBehaviour                   = 180
	AudioFilter                      = 181
	WindZone                         = 182
	Cloth                            = 183
	SubstanceArchive                 = 184
	ProceduralMaterial               = 185
	ProceduralTexture                = 186
	OffMeshLink                      = 191
	OcclusionArea                    = 192
	Tree                             = 193
	NavMesh                          = 194
	NavMeshAgent                     = 195
	NavMeshSettings                  = 196
	LightProbeCloud                  = 197
	ParticleSystem                   = 198
	ParticleSystemRenderer           = 199
	LODGroup                         = 205
	LightProbeGroup                  = 220
	Prefab                           = 1001
	EditorExtensionImpl              = 1002
	AssetImporter                    = 1003
	AssetDatabase                    = 1004
	Mesh3DSImporter                  = 1005
	TextureImporter                  = 1006
	ShaderImporter                   = 1007
	AudioImporter                    = 1020
	HierarchyState                   = 1026
	GUIDSerializer                   = 1027
	AssetMetaData                    = 1028
	DefaultAsset                     = 1029
	DefaultImporter                  = 1030
	TextScriptImporter               = 1031
	NativeFormatImporter             = 1034
	MonoImporter                     = 1035
	AssetServerCache                 = 1037
	LibraryAssetImporter             = 1038
	ModelImporter                    = 1040
	FBXImporter                      = 1041
	TrueTypeFontImporter             = 1042
	MovieImporter                    = 1044
	EditorBuildSettings              = 1045
	DDSImporter                      = 1046
	InspectorExpandedState           = 1048
	AnnotationManager                = 1049
	MonoAssemblyImporter             = 1050
	EditorUserBuildSettings          = 1051
	PVRImporter                      = 1052
	SubstanceImporter                = 1112
)

func (c ClassID) String() string {
	switch c {
	case Null:
		return "Null"
	case GameObject:
		return "GameObject"
	case Component:
		return "Component"
	case LevelGameManager:
		return "LevelGameManager"
	case Transform:
		return "Transform"
	case TimeManager:
		return "TimeManager"
	case GlobalGameManager:
		return "GlobalGameManager"
	case Behaviour:
		return "Behaviour"
	case GameManager:
		return "GameManager"
	case AudioManager:
		return "AudioManager"
	case ParticleAnimator:
		return "ParticleAnimator"
	case InputManager:
		return "InputManager"
	case EllipsoidParticleEmitter:
		return "EllipsoidParticleEmitter"
	case Pipeline:
		return "Pipeline"
	case EditorExtension:
		return "EditorExtension"
	case Camera:
		return "Camera"
	case Material:
		return "Material"
	case MeshRenderer:
		return "MeshRenderer"
	case Renderer:
		return "Renderer"
	case ParticleRenderer:
		return "ParticleRenderer"
	case Texture:
		return "Texture"
	case Texture2D:
		return "Texture2D"
	case Scene:
		return "Scene"
	case RenderManager:
		return "RenderManager"
	case MeshFilter:
		return "MeshFilter"
	case OcclusionPortal:
		return "OcclusionPortal"
	case Mesh:
		return "Mesh"
	case Skybox:
		return "Skybox"
	case QualitySettings:
		return "QualitySettings"
	case Shader:
		return "Shader"
	case TextAsset:
		return "TextAsset"
	case NotificationManager:
		return "NotificationManager"
	case Rigidbody:
		return "Rigidbody"
	case PhysicsManager:
		return "PhysicsManager"
	case Collider:
		return "Collider"
	case Joint:
		return "Joint"
	case HingeJoint:
		return "HingeJoint"
	case MeshCollider:
		return "MeshCollider"
	case BoxCollider:
		return "BoxCollider"
	case AnimationManager:
		return "AnimationManager"
	case AnimationClip:
		return "AnimationClip"
	case ConstantForce:
		return "ConstantForce"
	case WorldParticleCollider:
		return "WorldParticleCollider"
	case TagManager:
		return "TagManager"
	case AudioListener:
		return "AudioListener"
	case AudioSource:
		return "AudioSource"
	case AudioClip:
		return "AudioClip"
	case RenderTexture:
		return "RenderTexture"
	case MeshParticleEmitter:
		return "MeshParticleEmitter"
	case ParticleEmitter:
		return "ParticleEmitter"
	case Cubemap:
		return "Cubemap"
	case GUILayer:
		return "GUILayer"
	case ScriptMapper:
		return "ScriptMapper"
	case TrailRenderer:
		return "TrailRenderer"
	case DelayedCallManager:
		return "DelayedCallManager"
	case TextMesh:
		return "TextMesh"
	case RenderSettings:
		return "RenderSettings"
	case Light:
		return "Light"
	case CGProgram:
		return "CGProgram"
	case Animation:
		return "Animation"
	case MonoBehaviour:
		return "MonoBehaviour"
	case MonoScript:
		return "MonoScript"
	case MonoManager:
		return "MonoManager"
	case Texture3D:
		return "Texture3D"
	case Projector:
		return "Projector"
	case LineRenderer:
		return "LineRenderer"
	case Flare:
		return "Flare"
	case Halo:
		return "Halo"
	case LensFlare:
		return "LensFlare"
	case FlareLayer:
		return "FlareLayer"
	case HaloLayer:
		return "HaloLayer"
	case NavMeshLayers:
		return "NavMeshLayers"
	case HaloManager:
		return "HaloManager"
	case Font:
		return "Font"
	case PlayerSettings:
		return "PlayerSettings"
	case NamedObject:
		return "NamedObject"
	case GUITexture:
		return "GUITexture"
	case GUIText:
		return "GUIText"
	case GUIElement:
		return "GUIElement"
	case PhysicMaterial:
		return "PhysicMaterial"
	case SphereCollider:
		return "SphereCollider"
	case CapsuleCollider:
		return "CapsuleCollider"
	case SkinnedMeshRenderer:
		return "SkinnedMeshRenderer"
	case FixedJoint:
		return "FixedJoint"
	case RaycastCollider:
		return "RaycastCollider"
	case BuildSettings:
		return "BuildSettings"
	case AssetBundle:
		return "AssetBundle"
	case CharacterController:
		return "CharacterController"
	case CharacterJoint:
		return "CharacterJoint"
	case SpringJoint:
		return "SpringJoint"
	case WheelCollider:
		return "WheelCollider"
	case ResourceManager:
		return "ResourceManager"
	case NetworkView:
		return "NetworkView"
	case NetworkManager:
		return "NetworkManager"
	case PreloadData:
		return "PreloadData"
	case MovieTexture:
		return "MovieTexture"
	case ConfigurableJoint:
		return "ConfigurableJoint"
	case TerrainCollider:
		return "TerrainCollider"
	case MasterServerInterface:
		return "MasterServerInterface"
	case TerrainData:
		return "TerrainData"
	case LightmapSettings:
		return "LightmapSettings"
	case WebCamTexture:
		return "WebCamTexture"
	case EditorSettings:
		return "EditorSettings"
	case InteractiveCloth:
		return "InteractiveCloth"
	case ClothRenderer:
		return "ClothRenderer"
	case SkinnedCloth:
		return "SkinnedCloth"
	case AudioReverbFilter:
		return "AudioReverbFilter"
	case AudioHighPassFilter:
		return "AudioHighPassFilter"
	case AudioChorusFilter:
		return "AudioChorusFilter"
	case AudioReverbZone:
		return "AudioReverbZone"
	case AudioEchoFilter:
		return "AudioEchoFilter"
	case AudioLowPassFilter:
		return "AudioLowPassFilter"
	case AudioDistortionFilter:
		return "AudioDistortionFilter"
	case AudioBehaviour:
		return "AudioBehaviour"
	case AudioFilter:
		return "AudioFilter"
	case WindZone:
		return "WindZone"
	case Cloth:
		return "Cloth"
	case SubstanceArchive:
		return "SubstanceArchive"
	case ProceduralMaterial:
		return "ProceduralMaterial"
	case ProceduralTexture:
		return "ProceduralTexture"
	case OffMeshLink:
		return "OffMeshLink"
	case OcclusionArea:
		return "OcclusionArea"
	case Tree:
		return "Tree"
	case NavMesh:
		return "NavMesh"
	case NavMeshAgent:
		return "NavMeshAgent"
	case NavMeshSettings:
		return "NavMeshSettings"
	case LightProbeCloud:
		return "LightProbeCloud"
	case ParticleSystem:
		return "ParticleSystem"
	case ParticleSystemRenderer:
		return "ParticleSystemRenderer"
	case LODGroup:
		return "LODGroup"
	case LightProbeGroup:
		return "LightProbeGroup"
	case Prefab:
		return "Prefab"
	case EditorExtensionImpl:
		return "EditorExtensionImpl"
	case AssetImporter:
		return "AssetImporter"
	case AssetDatabase:
		return "AssetDatabase"
	case Mesh3DSImporter:
		return "Mesh3DSImporter"
	case TextureImporter:
		return "TextureImporter"
	case ShaderImporter:
		return "ShaderImporter"
	case AudioImporter:
		return "AudioImporter"
	case HierarchyState:
		return "HierarchyState"
	case GUIDSerializer:
		return "GUIDSerializer"
	case AssetMetaData:
		return "AssetMetaData"
	case DefaultAsset:
		return "DefaultAsset"
	case DefaultImporter:
		return "DefaultImporter"
	case TextScriptImporter:
		return "TextScriptImporter"
	case NativeFormatImporter:
		return "NativeFormatImporter"
	case MonoImporter:
		return "MonoImporter"
	case AssetServerCache:
		return "AssetServerCache"
	case LibraryAssetImporter:
		return "LibraryAssetImporter"
	case ModelImporter:
		return "ModelImporter"
	case FBXImporter:
		return "FBXImporter"
	case TrueTypeFontImporter:
		return "TrueTypeFontImporter"
	case MovieImporter:
		return "MovieImporter"
	case EditorBuildSettings:
		return "EditorBuildSettings"
	case DDSImporter:
		return "DDSImporter"
	case InspectorExpandedState:
		return "InspectorExpandedState"
	case AnnotationManager:
		return "AnnotationManager"
	case MonoAssemblyImporter:
		return "MonoAssemblyImporter"
	case EditorUserBuildSettings:
		return "EditorUserBuildSettings"
	case PVRImporter:
		return "PVRImporter"
	case SubstanceImporter:
		return "SubstanceImporter"
	default:
		return "UnExpected (New Version?)"
	}
}
