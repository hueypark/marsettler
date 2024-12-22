use avian2d::prelude::*;
use bevy::prelude::*;

pub fn spawn_snake_head(mut cmds: Commands) {
    const LENGTH: f32 = 5.0;

    cmds.spawn((
        SpriteBundle {
            sprite: Sprite {
                color: SNAKE_HEAD_COLOR,
                custom_size: Some(Vec2::splat(LENGTH)),
                ..default()
            },
            transform: Transform {
                scale: Vec3::splat(LENGTH),
                ..default()
            },
            ..default()
        },
        RigidBody::Dynamic,
        LinearDamping(1.6),
        AngularDamping(1.0),
        Collider::rectangle(LENGTH, LENGTH),
        SnakeHead {
            desired_position: Vec3::ZERO,
        },
    ));
}

pub fn rotate_snakes(mut query: Query<(&mut Transform, &mut ExternalTorque, &SnakeHead)>) {
    const TORQUE: f32 = 300_000.0;

    for (transform, mut et, snake_head) in query.iter_mut() {
        let direction = (snake_head.desired_position - transform.translation).normalize();
        let rot = transform.rotation * Vec3::Y;

        let cross = direction.cross(rot).z;

        let mul = if 0.0 <= cross { -1.0 } else { 1.0 };

        et.set_torque(TORQUE * mul);
    }
}

pub fn move_snakes(mut query: Query<(&LinearVelocity, &mut ExternalForce, &Transform)>) {
    const MAX_SPEED: f32 = 250.0;
    const FORCE: f32 = 300_000.0;

    for (vel, mut force, transform) in query.iter_mut() {
        let forward = (transform.rotation * Vec3::Y).truncate();

        let speed = vel.0.dot(forward);
        if MAX_SPEED <= speed {
            continue;
        }

        force.set_force(forward * FORCE);
    }
}

#[derive(Component)]
pub struct SnakeHead {
    pub desired_position: Vec3,
}

const SNAKE_HEAD_COLOR: Color = Color::srgb(0.7, 0.7, 0.7);
