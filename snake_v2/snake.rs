use avian2d::prelude::*;
use bevy::prelude::*;

pub fn spawn_snake_head(mut cmds: Commands) {
    const LENGTH: f32 = 10.0;

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
        Collider::rectangle(LENGTH, LENGTH),
        SnakeHead {
            desired_position: Vec3::ZERO,
        },
    ));
}

pub fn rotate_snakes(time: Res<Time>, mut query: Query<(&mut Transform, &SnakeHead)>) {
    const ROT_SPEED: f32 = 90.0;

    let ds = time.delta_seconds();
    let rot_speed = ROT_SPEED.to_radians() * ds;

    for (mut transform, snake_head) in query.iter_mut() {
        let desired_forward = snake_head.desired_position - transform.translation;
        if desired_forward.length() < 0.1 {
            continue;
        }

        let desired_rotation = Quat::from_rotation_arc(Vec3::Y, (desired_forward).normalize());

        transform.rotation = transform.rotation.slerp(desired_rotation, rot_speed);
    }
}

pub fn move_snakes(time: Res<Time>, mut query: Query<(&mut LinearVelocity, &Transform)>) {
    const ACC: f32 = 200.0;
    const MAX_SPEED: f32 = 150.0;

    let ds = time.delta_seconds();

    for (mut vel, transform) in query.iter_mut() {
        let forward = (transform.rotation * Vec3::Y).truncate();
        let desired_velocity = forward * MAX_SPEED;

        let diff = desired_velocity - vel.0;

        if diff.length() < ACC {
            vel.0 = desired_velocity;
        } else {
            vel.0 += diff.normalize() * ACC * ds;
        }
    }
}

#[derive(Component)]
pub struct SnakeHead {
    pub desired_position: Vec3,
}

const SNAKE_HEAD_COLOR: Color = Color::srgb(0.7, 0.7, 0.7);
