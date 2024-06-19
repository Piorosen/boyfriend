#%%
from pydub import AudioSegment
import os
#%%
def raw_to_mp3(raw_file_path, output_mp3_path, sample_width, frame_rate, channels):
    # Load raw audio data
    raw_audio = AudioSegment.from_raw(
        raw_file_path,
        sample_width=sample_width,
        frame_rate=frame_rate,
        channels=channels
    )
    
    # Export as MP3
    raw_audio.export(output_mp3_path, format="mp3")
    print(f"Conversion successful: {output_mp3_path}")
#%%
# Example usage
raw_file_path = "./cmake-build-debug/data.raw"
output_mp3_path = "./cmake-build-debug/data.mp3"
sample_width = 2  # 2 bytes for 16-bit audio
frame_rate = 44100  # 44.1 kHz
channels = 1  # Stereo

raw_to_mp3(raw_file_path, output_mp3_path, sample_width, frame_rate, channels)

# %%
