use bevy::prelude::*;

pub fn setup(mut cmds: Commands) {
    cmds.spawn(Camera2d::default());
}
