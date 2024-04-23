package whispercppgo

import "time"

// WhisperCPPOptions defines the command-line flags for the whisper.cpp tool.
type WhisperCPPOptions struct {
	Threads        int           // Number of threads to use during computation
	Processors     int           // Number of processors to use during computation
	OffsetT        int           // Time offset in milliseconds
	OffsetN        int           // Segment index offset
	Duration       time.Duration // Duration of audio to process
	MaxContext     int           // Maximum number of text context tokens to store
	MaxLen         int           // Maximum segment length in characters
	SplitOnWord    bool          // Split on word rather than on token
	BestOf         int           // Number of best candidates to keep
	BeamSize       int           // Beam size for beam search
	AudioCtx       int           // Audio context size (0 - all)
	WordThold      float64       // Word timestamp probability threshold
	EntropyThold   float64       // Entropy threshold for decoder fail
	LogprobThold   float64       // Log probability threshold for decoder fail
	DebugMode      bool          // Enable debug mode (eg. dump log_mel)
	Translate      bool          // Translate from source language to english
	Diarize        bool          // Enable stereo audio diarization
	TinyDiarize    bool          // Enable tinydiarize (requires a tdrz model)
	NoFallback     bool          // Do not use temperature fallback while decoding
	OutputTxt      bool          // Output result in a text file
	OutputVtt      bool          // Output result in a vtt file
	OutputSrt      bool          // Output result in a srt file
	OutputLrc      bool          // Output result in a lrc file
	OutputWords    bool          // Output script for generating karaoke video
	FontPath       string        // Path to a monospace font for karaoke video
	OutputCsv      bool          // Output result in a CSV file
	OutputJson     bool          // Output result in a JSON file
	OutputJsonFull bool          // Include more information in the JSON file
	OutputFile     string        // Output file path (without file extension)
	NoPrints       bool          // Do not print anything other than the results
	PrintSpecial   bool          // Print special tokens
	PrintColors    bool          // Print colors
	PrintProgress  bool          // Print progress
	NoTimestamps   bool          // Do not print timestamps
	Language       string        // Spoken language ('auto' for auto-detect)
	DetectLanguage bool          // Exit after automatically detecting language
	Prompt         string        // Initial prompt (max n_text_ctx/2 tokens)
	Model          string        // Model path
	File           string        // Input WAV file path
	OvEDevice      string        // The OpenVINO device used for encode inference
	DtwModel       string        // Compute token-level timestamps
	LogScore       bool          // Log best decoder scores of tokens
	NoGPU          bool          // Disable GPU
	SuppressRegex  string        // Regular expression matching tokens to suppress
	Grammar        string        // GBNF grammar to guide decoding
	GrammarRule    string        // Top-level GBNF grammar rule name
	GrammarPenalty float64       // Scales down logits of nongrammar tokens
}
