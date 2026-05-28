This program is the solution I wrote for gopphercises quiz game parts 1 and 2: https://gophercises.com/ 

The below readme.md file description was generated with the help of Gemini AI. However, the code I wrote was _not_ AI generated, I wrote it myself. 

# Go Quiz Game

A customizable, terminal based quiz application written in Go. The program reads quiz questions and answers from a CSV file, tracks user scores, features a strict concurrent countdown timer, and offers automated question shuffling.

## Features

* **Polite Greeting**: Welcomes the user before starting the quiz.
* **CSV Quiz Parsing**: Dynamically reads questions and answers from structured CSV files.
* **Score Tracking**: Evaluates user input immediately and provides a final tally of correct answers vs. total questions.
* **Concurrent Timer**: Runs an asynchronous countdown timer that instantly stops the game when time expires.
* **Robust Input Sanitization**: Uses `strings.TrimSpace` to ensure extra whitespace around answers doesn't cause false negatives.
* **Quiz Shuffling**: Randomizes the order of questions to ensure a unique experience on every run.

---

## Configuration Flags

Customizes the behavior of the quiz directly from the terminal using the following built-in flags:


| Flag | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `-csv` | `string` | `"problems.csv"` | Path to the custom quiz CSV file. |
| `-timeout` | `duration` | `30s` | The quiz time limit (e.g., `30s`, `1m`). |
| `-shuffle` | `bool` | `false` | Set to `true` to randomize the question order. |

---

## CSV File Format

My quiz data files follow a simple two column structure with no headers:

```csv
5+5,10
7*3,21
```
* **Column 1**: The question prompt.
* **Column 2**: The exact expected answer.

---

## Gameplay Mechanics

1. **Immediate Start**: The application displays a greeting and instantly starts the timer along with the first question.
2. **Immediate Progression**: The quiz progresses to the next question immediately after an input, regardless of whether the answer was right or wrong.
3. **Strict Timeout**: The program forcibly terminates the quiz the exact second the timer expires, even if you are currently typing an answer.
4. **Scoring Logic**: Invalid inputs, formatting anomalies, and unanswered questions due to timeout are automatically marked as incorrect.


