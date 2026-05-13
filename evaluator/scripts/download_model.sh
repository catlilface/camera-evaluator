curl -L -o .model/MUSIQ.tar.gz https://www.kaggle.com/api/v1/models/google/musiq/tensorFlow2/koniq-10k/1/download
tar -xzvf .model/MUSIQ.tar.gz -C .model
rm -f .model/MUSIQ.tar.gz
