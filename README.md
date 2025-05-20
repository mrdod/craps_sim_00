# Craps Simulator

A web-based Craps game simulator written in Go. This application provides a simple and intuitive interface to play the classic casino dice game.

## Features

- Realistic Craps game rules
- Modern web interface
- Real-time game state updates
- Responsive design

## Prerequisites

- Go 1.21 or later

## Running the Game

1. Clone this repository
2. Navigate to the project directory
3. Run the following command:
   ```bash
   go run .
   ```
4. Open your web browser and visit `http://localhost:8080`

## How to Play

1. Click the "Roll Dice" button to start the game
2. On the first roll:
   - Rolling 7 or 11 is a "Natural" - you win!
   - Rolling 2, 3, or 12 is "Craps" - you lose!
   - Any other number becomes your "Point"
3. After the point is set:
   - Roll the point again to win
   - Roll a 7 to lose
   - Any other roll continues the game
4. Click "New Game" to start over

## Project Structure

- `main.go` - Web server and request handlers
- `game.go` - Core game logic
- `templates/index.html` - Web interface 