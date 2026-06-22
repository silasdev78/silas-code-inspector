package learner

import (
	"encoding/json"
	"os"
	"sync"
)

type PatternWeight struct {
	PatternID string  `json:"pattern_id"`
	BaseScore float64 `json:"base_score"`
	Feedback  float64 `json:"feedback"`
}

type State struct {
	Weights map[string]PatternWeight `json:"weights"`
}

type Learner struct {
	mu    sync.RWMutex
	state State
	path  string
}

func NewLearner(path string) (*Learner, error) {
	l := &Learner{
		state: State{Weights: make(map[string]PatternWeight)},
		path:  path,
	}
	data, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(data, &l.state)
	}
	return l, nil
}

func (l *Learner) Save() error {
	l.mu.RLock()
	defer l.mu.RUnlock()
	data, err := json.MarshalIndent(l.state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.path, data, 0644)
}

func (l *Learner) GetWeight(patternID string) float64 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if w, ok := l.state.Weights[patternID]; ok {
		return w.BaseScore + w.Feedback
	}
	return 1.0
}

func (l *Learner) Adjust(patternID string, positive bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	w, ok := l.state.Weights[patternID]
	if !ok {
		w = PatternWeight{PatternID: patternID, BaseScore: 1.0}
	}
	if positive {
		w.Feedback += 0.1
	} else {
		w.Feedback -= 0.1
	}
	l.state.Weights[patternID] = w
}
