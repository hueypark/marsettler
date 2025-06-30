use bevy::prelude::*;
use bevy_hanabi::prelude::*;
use rand::Rng;
use std::time::Instant;

#[derive(Resource)]
pub struct EffectResource {
    handle: Handle<EffectAsset>,
    lifetime: f32,
}

pub fn setup_blacksmith_hammer_sparks(
    mut cmds: Commands,
    mut effects: ResMut<Assets<EffectAsset>>,
) {
    const LIFETIME: f32 = 3.0;
    const PARTICLE_COUNT: u32 = 100;

    let writer = ExprWriter::new();

    let mut gradient = Gradient::new();
    gradient.add_key(0.0, Vec4::new(1., 1., 1., 1.));
    gradient.add_key(1.0, Vec4::splat(0.));

    let init_pos = SetPositionSphereModifier {
        center: writer.lit(Vec3::new(0.0, 1.0, 0.0)).expr(),
        radius: writer.lit(1.).expr(),
        dimension: ShapeDimension::Surface,
    };

    let init_vel = SetVelocitySphereModifier {
        center: writer.lit(Vec3::ZERO).expr(),
        speed: writer.lit(1.0).uniform(writer.lit(30.0)).expr(),
    };

    let lifetime = writer.lit(LIFETIME).expr();
    let init_lifetime = SetAttributeModifier::new(Attribute::LIFETIME, lifetime);

    let accel = writer.lit(Vec3::new(0., -9.81, 0.)).expr();
    let update_accel = AccelModifier::new(accel);

    let effect = EffectAsset::new(
        PARTICLE_COUNT,
        SpawnerSettings::once((PARTICLE_COUNT as f32).into()),
        writer.finish(),
    )
    .init(init_pos)
    .init(init_vel)
    .init(init_lifetime)
    .update(update_accel)
    .render(ColorOverLifetimeModifier {
        gradient,
        ..default()
    });

    cmds.insert_resource(EffectResource {
        handle: effects.add(effect),
        lifetime: LIFETIME,
    });
}

#[derive(Component)]
pub struct Effect {
    despawn_time: Instant,
}

pub fn update_spawn(
    mut cmds: Commands,
    keys: Res<ButtonInput<KeyCode>>,
    effect: Res<EffectResource>,
) {
    if keys.just_pressed(KeyCode::Space) {
        cmds.spawn((
            Effect {
                despawn_time: Instant::now() + std::time::Duration::from_secs_f32(effect.lifetime),
            },
            ParticleEffect {
                handle: effect.handle.clone(),
                prng_seed: Some(rand::rng().random()),
            },
            Transform::from_translation(Vec3::Y),
        ));
    }
}

pub fn update(mut cmds: Commands, mut query: Query<(Entity, &Effect)>) {
    for (entity, effect) in query.iter_mut() {
        if Instant::now() > effect.despawn_time {
            cmds.entity(entity).despawn();
        }
    }
}
