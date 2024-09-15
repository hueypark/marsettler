use crate::status;
use bevy::prelude::*;

#[derive(Component)]
pub struct StatusUI;

pub fn update_status(st: Res<status::Status>, mut query: Query<&mut Text, With<StatusUI>>) {
    for mut text in query.iter_mut() {
        text.sections[0].value = format!("Time: {:.2}s", st.time());
    }
}
