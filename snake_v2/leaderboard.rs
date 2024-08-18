use bevy::prelude::*;
use bevy_jornet::Leaderboard;

pub fn setup_leaderboard(mut leaderboard: ResMut<Leaderboard>) {
    leaderboard.create_player(None);
    leaderboard.refresh_leaderboard();
}

#[derive(Event)]
pub struct SendScoreEvent {
    pub score: f32,
}

pub fn send_score(mut event_reader: EventReader<SendScoreEvent>, leaderboard: ResMut<Leaderboard>) {
    for event in event_reader.read() {
        debug!("send_score: {}", event.score);

        leaderboard.send_score(event.score);
        leaderboard.refresh_leaderboard();
    }
}
