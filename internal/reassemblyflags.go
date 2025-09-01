package internal

type FeatureE uint64
type ECommandFlags uint64
type EExplosive uint64

// All block features
// https://www.anisopteragames.com/docs/
const (
	initFeatureFlag        FeatureE = (1 << iota) >> 1
	FeatureCommand                  // Designates a command block ( if this dies, ship dies).
	FeatureThruster                 // Thruster, provides thrust in a direction.
	FeatureGenerator                // Provides energy for weapons.
	FeaturePerishable               // has limited lifetime.
	FeatureTurret                   // Turreted weapon - use with CANNON or LASER.
	FeatureLaunch                   // Used internally for missiles/seeds that can be launched by player. #DO NOT USE, NOT NECESSARY.#
	FeatureCannon                   // Shoots projectiles. ( can be fragmentation )
	FeatureLaser                    // Shoots beams. ( can be healing beams )
	FeatureAutofire                 // Weapon automatically targets and fires at enemies.
	FeatureShield                   // Round shield, blocks enemy weapons.
	FeatureTorquer                  // Provides torque. (rotation)
	FeatureLauncher                 // Generates launchable objects. ( Missiles, mines, torpedoes, etc) and launches them on player input.
	FeatureExplode                  // Explodes on contact with enemy, or expiration.
	FeatureAssembler                // Allows cluster to reassemble ( regrow ), and use tractored parts to rebuild.
	FeatureRegrower                 // Allows cluster to reassemble ( regrow ), but can not tractor new parts to help rebuild.
	FeatureCannonBoost              // Attaches to cannon blocks and increases their power.
	FeatureInvulnerable             // Block is indestructible.
	FeatureNoregen                  // Does not regenerate health.
	FeaturePersistent               // Used internally for map objectives - stations, agents. #Not Recommended Ever#
	FeatureEnvironmental            // Allows plants/seeds/roots to attach to it.
	FeatureTractor                  // Can harvest resources.
	FeatureRoot                     // Can attach to asteroids/ENVIRONMENTAL.
	FeatureGrow                     // Currently growing. #Not Recommended Ever#
	FeaturePhotosynth               // Generates Resources.
	FeatureAutolaunch               // Launches automatically, without being triggered.
	FeatureFreeres                  // Does not use resources or drop resources. ( good for drones / missiles )
	FeatureFactory                  // Can spawn child blocks.
	FeatureSeed                     // Can plant itself on asteroids/ENVIRONMENTAL, will thruster if also THRUSTER.
	FeatureMelee                    // Does damage on contact with enemy.
	FeatureUngrow                   // Currently ungrowing. ( Not useful practically )
	FeatureUnique                   // Special feature for use in editor - can't be duplicated.
	FeatureCharging                 // Charging weapon - use with CANNON or LASER.
	FeatureTransient                // Used internally to mark blocks only attached temporarily, not part of blueprint.
	FeatureSelffactory              // Can spawn own design.
	FeatureNoclip                   // No collisions.
	FeatureInvisible                // Don't render block.
	FeatureBumper                   // Disables collision damage.
	FeatureTeleporter               // Can teleport ship.
	FeatureDeactivates              // Deactivates instead of being destroyed. (like stations)
	FeatureTelespawn                // Like FACTORY but spawns ships by teleporting them in, already grown.
	FeatureNoclipAlly               // Blocks with this flag will not collide with friendly blocks.
	FeatureIntlines                 // Draws the lines between blocks.
	FeatureOneuse                   // Block will act once, then ungrow. ( Cannon, Laser, drone )
	FeatureNorecolor                // Will not be recoloured when the player chooses their colours in the spawn screen.
	FeatureNopalette                // Will not appear in the player's database at all.
	FeatureLauncherBarrage          // Will grow all launchables and fire them at once.
	FeatureAlwaysfire               // Will always fire, and be ignored by the AI.
	FeaturePalette                  // Will appear in the player's database without needing to be on a ship.
	FeatureNeverfire                // Will never fire, and be ignored by the AI.
)

const (
	CommandFlagNone             ECommandFlags = (1 << iota) >> 1
	CommandFlagMetamorphosis                  // Ai will occasionally change blueprints.
	CommandFlagFollower                       // Follows player.
	CommandFlagAttack                         // Special flag for tournament mode, attack ruthlessly.
	CommandFlagFlocking                       // Align with nearby allies.
	CommandFlagReckless                       // Run away less.
	CommandFlagAggressive                     // Initiate attack more easily.
	CommandFlagCautious                       // Initiate attack less easily.
	CommandFlagSocial                         // Call for help when attacked.
	CommandFlagPeaceful                       // Never initiate attack.
	CommandFlagWander                         // Wander randomly if nothing else to do. ( recommended )
	CommandFlagHatesPlants                    // Kill plants if in range.
	CommandFlagForgiving                      // Stop attacking more easily.
	CommandFlagTractorTransient               // Grab blocks from the environment and use them.
	CommandFlagDodges                         // Dodge projectiles.
	CommandFlagRippleFire                     // Use ripple fire on weapons.
	CommandFlagSpreadFire                     // Use spread fire on weapons.
	CommandFlagBadAim                         // Aim poorly.
	CommandFlagPointDefense                   // Act like a point defense drone.
	CommandFlagInactive                       // Become a vegetable.
	CommandFlagSmartFire                      // Use spread fire when enemy is expected to dodge.
	CommandFlagNoParent                       // Don't follow leader.
	CommandFlagChildrenSet                    // Don't change children - set by player.
	CommandFlagBlueprintSet                   // Don't change blueprint - set by player.
	CommandFlagHangout                        // Used to make the Anisoptera ship spin in circles. ;-)
	CommandFlagPacifist                       // Never attack even when attacked.
	CommandFlagAlwaysKite                     // Always attack from max range.
	CommandFlagAlwaysRush                     // Always attack at closest range regardless of incoming damage.
	CommandFlagAlwaysManeuver                 // Always dodging while attacking.
	CommandFlagFireAtWill                     // Ignore parent ship target.
)

const (
	initExplosiveFlag      EExplosive = (1 << iota) >> 1
	ExplosiveEnabled                  // Will explode.
	ExplosiveProximity                // Will snap it's explosion to the closest object relative to explode radius.
	ExplosiveFinal                    // Will explode at the end of it's range / when the bullet's damage output is exhausted.

	ExplosiveFragProximity            // Fragments when within proximity to an object.
	ExplosiveFragFinal                // Fragments at the end of it's range / when the bullet's damage output is exhausted.
	ExplosiveFragImpact               // Fragments upon contact with an object.
)
